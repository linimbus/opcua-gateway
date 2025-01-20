package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/astaxie/beego/logs"
)

type OrderType int

const (
	ORDER_UP OrderType = iota
	ORDER_DOWN
	ORDER_TOP
	ORDER_BOTTOM
)

func ClientTest(addr string) error {
	cli, err := NewClient(addr)
	if err != nil {
		logs.Error("connect [%s] failed: %s", addr, err.Error())
		return err
	}
	defer cli.Close()
	return nil
}

type OpcuaClientData struct {
	name   string
	nodes  []NodeInfo
	values []*NodeValue
}

type OpcuaStoreData struct {
	table  string
	values []string
}

type OpcuaBasket struct {
	baskedID    string
	siliconData []*NodeValue
}

type OpcuaServer struct {
	sync.RWMutex
	sync.WaitGroup

	cfg      Config
	shutdown bool

	db     *DataSave
	dbChan chan interface{}

	basketCache map[string]OpcuaBasket

	server      *Server
	serverChan  chan interface{}
	serverCache map[string]NodeInfo

	clients map[string]*Client
	stats   map[string]*StatItem
}

func (opc *OpcuaServer) dataStoreTask() {
	defer opc.Done()

	logs.Info("data store task startup")

	stat := opc.stats[STAT_MYSQL]

	err := opc.db.TableExpired(true)
	if err != nil {
		logs.Warning("table expired enable failed, %s", err.Error())
		atomic.AddUint64(&stat.OperFail, 1)
	}

	for {
		data := <-opc.dbChan

		if _, ok := data.(bool); ok {
			break
		}

		storeData, ok := data.(OpcuaStoreData)
		if ok {
			err := opc.db.TableWrite(storeData.table, storeData.values)
			if err != nil {
				logs.Warning("data store table write %s failed, %s", storeData.table, err.Error())
				atomic.AddUint64(&stat.OperFail, 1)
			} else {
				atomic.AddUint64(&stat.OperOK, 1)
			}
		}
	}

	err = opc.db.TableExpired(false)
	if err != nil {
		logs.Warning("table expired disable failed!, %s", err.Error())
		atomic.AddUint64(&stat.OperFail, 1)
	}

	logs.Info("data store task shutdown")
}

func (opc *OpcuaServer) serverTask() {
	defer opc.Done()

	logs.Info("server sync data task startup")

	for {
		data := <-opc.serverChan

		if _, ok := data.(bool); ok {
			break
		}

		nodeData, ok := data.(OpcuaClientData)
		if !ok {
			continue
		}

		success := true
		for index, node := range nodeData.nodes {
			serverName := fmt.Sprintf("%s.%s", nodeData.name, node.NodeID)
			serverNode, b := opc.serverCache[serverName]
			if !b {
				continue
			}
			err := opc.server.WriteNode(serverNode, *nodeData.values[index])
			if err != nil {
				success = false
				logs.Error("server write node %s failed, %s", serverNode.ToString(), err.Error())
			}
		}

		if success {
			atomic.AddUint64(&opc.stats[STAT_SERVER].OperOK, 1)
		} else {
			atomic.AddUint64(&opc.stats[STAT_SERVER].OperFail, 1)
		}
	}

	logs.Info("server sync data task shutdown")
}

func (opc *OpcuaServer) clientValue(query NodeInfo, nodes []NodeInfo, values []*NodeValue) *NodeValue {
	for i := 0; i < len(nodes); i++ {
		if query.Compare(nodes[i]) {
			return values[i].Clone()
		}
	}
	logs.Error("opcua client query node %s value not exist", query.ToString())
	return nil
}

func (opc *OpcuaServer) clientTask(cli *Client, name string) {
	defer opc.Done()

	logs.Info("opcua client %s startup", name)

	cfg := opc.cfg.ClientConfig(name)
	stat := opc.stats[STAT_CLIENT]

	nodeList := make([]NodeInfo, 0)
	for _, node := range cfg.NodeList {
		nodeList = append(nodeList, NodeInfo{
			NsIndex: node.NsIndex,
			NodeID:  node.NodeID,
		})
	}
	tableName := EscapeString(cfg.Name)

	for {
		time.Sleep(time.Duration(cfg.Timeout) * time.Millisecond)

		if opc.shutdown {
			break
		}

		if len(nodeList) == 0 {
			continue
		}

		nodeValues, err := cli.ReadNodes(nodeList)
		if err != nil {
			logs.Error("opcua client read nodes failed, %s", err.Error())
			cli.CheckState()
			atomic.AddUint64(&stat.OperFail, 1)
			continue
		}

		if cfg.Store && opc.db != nil {
			strList := make([]string, 0)
			for _, value := range nodeValues {
				strList = append(strList, value.ToString())
			}
			opc.dbChan <- OpcuaStoreData{
				table:  tableName,
				values: strList,
			}
			atomic.AddUint64(&stat.OperOK, 1)
		}

		if opc.server != nil {
			opc.serverChan <- OpcuaClientData{
				name:   name,
				nodes:  nodeList,
				values: nodeValues,
			}
			atomic.AddUint64(&stat.OperOK, 1)
		}
	}

	logs.Info("opcua client %s shutdown", cfg.Name)
}

func (opc *OpcuaServer) serverInit(cli *Client, name string) error {
	index, err := opc.server.AddNameSpace(name)
	if err != nil {
		logs.Error("opcua server add namespace %s failed, %s", name, err.Error())
		return err
	}

	for _, node := range opc.cfg.Server.NodeList {
		if node.ClientName != name {
			continue
		}

		value, err := cli.ReadNode(node.ClientNode)
		if err != nil {
			logs.Error("opcua server read node %s failed, %s", node.ClientNode.ToString(), err.Error())
			return err
		}

		clientNode := NodeInfo{NsIndex: uint32(index), NodeID: node.ClientName}
		serverNode := NodeInfo{NsIndex: uint32(index), NodeID: node.ServerName}

		err = opc.server.AddNode(clientNode, serverNode, node.ServerName, *value)
		if err != nil {
			logs.Error("opcua server add node %s failed, %s", serverNode.ToString(), err.Error())
			return err
		}

		opc.serverCache[node.ServerName] = serverNode

		logs.Info("opcua server add node %s success", serverNode.ToString())
	}

	return nil
}

func (opc *OpcuaServer) Close() {
	logs.Info("opcua server ready close")

	opc.shutdown = true
	opc.serverChan <- true
	opc.dbChan <- true
	opc.Wait()

	for _, cli := range opc.clients {
		cli.Close()
	}

	if opc.db != nil {
		opc.db.Close()
	}

	if opc.server != nil {
		opc.server.Close()
	}

	for _, stat := range opc.stats {
		stat.Clear()
	}

	logs.Info("opcua server close done")
}

func (opc *OpcuaServer) dataInit(db *DataSave) error {
	for _, cfg := range opc.cfg.Clients {
		if !cfg.Store || !cfg.Enable {
			continue
		}
		columns := make([]ColumnInfo, 0)
		for _, node := range cfg.NodeList {
			columns = append(columns, ColumnInfo{
				Name:    ColumnName(node.NodeID),
				Comment: EscapeString(node.NodeID),
			})
		}
		err := db.TableInit(EscapeString(cfg.Name), columns)
		if err != nil {
			logs.Error("opcua client table init %s failed", EscapeString(cfg.Name))
			return err
		}
	}

	logs.Info("opcua client table init success")
	return nil
}

func NewOpcuaServer(config Config, stats []*StatItem) (*OpcuaServer, error) {
	var err error

	opc := &OpcuaServer{
		cfg:         config,
		shutdown:    false,
		dbChan:      make(chan interface{}, 1024),
		serverChan:  make(chan interface{}, 1024),
		stats:       make(map[string]*StatItem),
		clients:     make(map[string]*Client),
		serverCache: make(map[string]NodeInfo),
		basketCache: make(map[string]OpcuaBasket),
	}

	defer func() {
		if err != nil {
			opc.Close()
		}
	}()

	for _, stat := range stats {
		opc.stats[stat.Name] = stat
	}

	if config.Datastore.Enable {
		opc.db, err = NewDataSave(config.Datastore)
		if err != nil {
			logs.Error("opcua client datastore init failed, %s", err.Error())
			return nil, err
		}
		err = opc.dataInit(opc.db)
		if err != nil {
			logs.Error("opcua client table init failed, %s", err.Error())
			return nil, err
		}
		opc.Add(1)
		go opc.dataStoreTask()
		opc.stats[STAT_MYSQL].Status = true
	}

	for _, cfg := range opc.cfg.Clients {
		if cfg.Enable {
			var cli *Client
			cli, err = NewClient(cfg.Endpoint)
			if err != nil {
				logs.Error("opcua client %s connect failed, %s", cfg.Name, err.Error())
				return nil, err
			}
			opc.clients[cfg.Name] = cli
		}
	}

	for name, cli := range opc.clients {
		opc.Add(1)
		go opc.clientTask(cli, name)
		opc.stats[STAT_CLIENT].Status = true
	}

	if config.Server.Enable {
		opc.server, err = NewServer(config.Server.Endpoint, config.Server.Port)
		if err != nil {
			logs.Error("opcua server init failed, %s", err.Error())
			return nil, err
		}
		for name, cli := range opc.clients {
			err = opc.serverInit(cli, name)
			if err != nil {
				logs.Error("opcua server init failed, %s", err.Error())
				return nil, err
			}
		}
		opc.Add(1)
		go opc.serverTask()
		opc.stats[STAT_SERVER].Status = true
	}

	return opc, nil
}
