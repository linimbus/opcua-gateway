package main

import (
	"sort"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type FromNodeItem struct {
	Index int

	name     string
	endpoint string
	node     NodeInfo

	checked bool
}

type FromNodeTable struct {
	sync.RWMutex

	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder

	items []*FromNodeItem
}

func (n *FromNodeTable) RowCount() int {
	return len(n.items)
}

func (n *FromNodeTable) Value(row, col int) interface{} {
	item := n.items[row]
	switch col {
	case 0:
		return item.Index
	case 1:
		return item.name
	case 2:
		return item.node.ToString()
	}
	panic("unexpected col")
}

func (n *FromNodeTable) Checked(row int) bool {
	return n.items[row].checked
}

func (n *FromNodeTable) SetChecked(row int, checked bool) error {
	n.items[row].checked = checked
	return nil
}

func (m *FromNodeTable) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order
	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]
		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}
			return !ls
		}
		switch m.sortColumn {
		case 0:
			return c(a.Index < b.Index)
		case 1:
			return c(a.name < b.name)
		case 2:
			return c(a.node.NodeID < b.node.NodeID)
		case 3:
			return c(a.node.ToString() < b.node.ToString())
		}
		panic("unreachable")
	})
	return m.SorterBase.Sort(col, order)
}

func (m *FromNodeTable) SelectItems() []*FromNodeItem {
	items := make([]*FromNodeItem, 0)
	for _, node := range m.items {
		if node.checked {
			items = append(items, node)
		}
	}
	return items
}

type ServerNodeItem struct {
	Index int

	clientName string
	clientNode NodeInfo
	serverName string
	serverNode NodeInfo

	checked bool
}

type ServerNodeTable struct {
	sync.RWMutex

	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder

	items []*ServerNodeItem
}

func (n *ServerNodeTable) RowCount() int {
	return len(n.items)
}

func (n *ServerNodeTable) Value(row, col int) interface{} {
	item := n.items[row]
	switch col {
	case 0:
		return item.Index
	case 1:
		return item.clientName
	case 2:
		return item.clientNode.ToString()
	case 3:
		return item.serverNode.ToString()
	}
	panic("unexpected col")
}

func (n *ServerNodeTable) Checked(row int) bool {
	return n.items[row].checked
}

func (n *ServerNodeTable) SetChecked(row int, checked bool) error {
	n.items[row].checked = checked
	return nil
}

func (m *ServerNodeTable) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order
	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]
		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}
			return !ls
		}
		switch m.sortColumn {
		case 0:
			return c(a.Index < b.Index)
		case 1:
			return c(a.clientName < b.clientName)
		case 2:
			return c(a.clientNode.ToString() < b.clientNode.ToString())
		case 3:
			return c(a.serverNode.ToString() < b.serverNode.ToString())
		}
		panic("unreachable")
	})
	return m.SorterBase.Sort(col, order)
}

var fromNodeTable *FromNodeTable
var serverNodeTable *ServerNodeTable

func init() {
	fromNodeTable = new(FromNodeTable)
	fromNodeTable.items = make([]*FromNodeItem, 0)

	serverNodeTable = new(ServerNodeTable)
	serverNodeTable.items = make([]*ServerNodeItem, 0)
}

func FromNodeTableClean() {
	fromNodeTable.Lock()
	defer fromNodeTable.Unlock()

	fromNodeTable.items = make([]*FromNodeItem, 0)
	fromNodeTable.PublishRowsReset()
	fromNodeTable.Sort(fromNodeTable.sortColumn, fromNodeTable.sortOrder)
}

func FromNodeTableInit(filter string, clients []ClientConfig) {
	fromNodeTable.Lock()
	defer fromNodeTable.Unlock()

	var index int
	items := make([]*FromNodeItem, 0)
	for _, client := range clients {
		if filter != "" && client.Name != filter {
			continue
		}
		for _, node := range client.NodeList {
			items = append(items, &FromNodeItem{
				Index:    index,
				name:     client.Name,
				endpoint: client.Endpoint,
				node:     node,
			})
			index++
		}
	}
	fromNodeTable.items = items
	fromNodeTable.PublishRowsReset()
	fromNodeTable.Sort(fromNodeTable.sortColumn, fromNodeTable.sortOrder)
}

func FromNodeTableSelect(flag bool) {
	fromNodeTable.Lock()
	defer fromNodeTable.Unlock()

	for _, item := range fromNodeTable.items {
		item.checked = flag
	}
	fromNodeTable.PublishRowsReset()
	fromNodeTable.Sort(fromNodeTable.sortColumn, fromNodeTable.sortOrder)
}

func serverNodeTableInit(server *ServerConfig) {
	items := make([]*ServerNodeItem, 0)
	for i, node := range server.NodeList {
		items = append(items, &ServerNodeItem{
			Index:      i,
			clientName: node.ClientName,
			clientNode: node.ClientNode,
			serverName: node.ServerName,
			serverNode: node.ServerNode,
		})
	}

	serverNodeTable.items = items
	serverNodeTable.PublishRowsReset()
	serverNodeTable.Sort(serverNodeTable.sortColumn, serverNodeTable.sortOrder)
}

func ServerNodeTableSelect(flag bool) {
	serverNodeTable.Lock()
	defer serverNodeTable.Unlock()

	for _, item := range serverNodeTable.items {
		item.checked = flag
	}

	serverNodeTable.PublishRowsReset()
	serverNodeTable.Sort(serverNodeTable.sortColumn, serverNodeTable.sortOrder)
}

func ServerNodeTableAdd(server *ServerConfig, nodes []*FromNodeItem) {
	serverNodeTable.Lock()
	defer serverNodeTable.Unlock()

	for _, node := range nodes {
		server.Add(node.name, node.endpoint, node.node)
	}

	serverNodeTableInit(server)
}

func ServerNodeTableUpdate(server *ServerConfig, node *ServerNodeItem) {
	serverNodeTable.Lock()
	defer serverNodeTable.Unlock()

	server.Update(node.serverName, node.serverNode)

	serverNodeTableInit(server)
}

func ServerNodeTableDelete(server *ServerConfig) {
	serverNodeTable.Lock()
	defer serverNodeTable.Unlock()

	for _, item := range serverNodeTable.items {
		if item.checked {
			server.Delete(item.serverName)
		}
	}

	serverNodeTableInit(server)
}

func ServerNodeTableClean(server *ServerConfig) {
	serverNodeTable.Lock()
	defer serverNodeTable.Unlock()

	server.Clean()
	serverNodeTableInit(server)
}

func ServerNodeTableInit(server *ServerConfig) {
	serverNodeTable.Lock()
	defer serverNodeTable.Unlock()

	serverNodeTableInit(server)
}

func ServerStartupTest(config ServerConfig) (*Server, error) {
	server, err := NewServer(config.Endpoint, config.Port)
	if err != nil {
		logs.Error("opcua server init failed, %s", err.Error())
		return nil, err
	}

	clients := make(map[string]*Client)

	defer func() {
		for _, client := range clients {
			client.Close()
		}
	}()

	for _, node := range config.NodeList {
		if _, ok := clients[node.ClientName]; !ok {
			cli, err := NewClient(node.ClientEndpoint)
			if err != nil {
				logs.Error("opcua client init failed, %s", err.Error())
				return nil, err
			}
			clients[node.ClientName] = cli
		}

		index, err := server.AddNameSpace(node.ClientName)
		if err != nil {
			logs.Error("opcua server add namespace %s failed, %s", node.ClientName, err.Error())
			return nil, err
		}

		cli := clients[node.ClientName]
		value, err := cli.ReadNode(node.ClientNode)
		if err != nil {
			logs.Error("opcua client read node %s failed, %s", node.ClientNode.ToString(), err.Error())
			return nil, err
		}

		clientNode := NodeInfo{NsIndex: index, NodeID: node.ClientName}
		serverNode := NodeInfo{NsIndex: index, NodeID: node.ServerName}

		err = server.AddNode(clientNode, serverNode, node.ServerName, *value)
		if err != nil {
			logs.Error("opcua server add node %s failed, %s", serverNode.ToString(), err.Error())
			return nil, err
		}

		logs.Info("opcua server add node %s success", serverNode.ToString())
	}

	return server, nil
}

func ServerNodeEditDialog(from walk.Form, config *ServerConfig, item *ServerNodeItem) {
	var dlg *walk.Dialog
	var acceptPB, cancelPB *walk.PushButton

	_, err := Dialog{
		AssignTo:      &dlg,
		Title:         "Edit Node Dialog",
		Icon:          walk.IconInformation(),
		MinSize:       Size{Width: 200, Height: 250},
		Size:          Size{Width: 200, Height: 250},
		Font:          DefaultFont(),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "Node Name:",
					},
					LineEdit{
						Text:        item.serverName,
						ToolTipText: item.serverName,
						Enabled:     false,
					},
					Label{
						Text: "Node Tag:",
					},
					LineEdit{
						Text:        item.serverNode.ToString(),
						ToolTipText: item.serverNode.ToString(),
						Enabled:     false,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "Accept",
						OnClicked: func() {
							logs.Info("server node single edit dialog accept")

							ServerNodeTableUpdate(config, item)
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo: &cancelPB,
						Text:     "Cancel",
						OnClicked: func() {
							logs.Info("server node single edit dialog cancel")

							dlg.Cancel()
						},
					},
					HSpacer{},
				},
			},
		},
	}.Run(from)

	if err != nil {
		logs.Error("ServerNodeEditDialog: %s", err.Error())
	}
}

func ServerEditDialog(from walk.Form, config *Config) {
	var dlg *walk.Dialog
	var nameLine *walk.LineEdit
	var acceptPB, cancelPB, startTest, stopTest *walk.PushButton
	var portNum *walk.NumberEdit
	var listenBox, clientBox *walk.ComboBox
	var fromNodeTableView *walk.TableView
	var toNodeTableView *walk.TableView
	var enableCB, fromCheckBox, serverCheckBox *walk.CheckBox
	var server *Server

	defer func() {
		if server != nil {
			server.Close()
			server = nil
		}
	}()

	interfaces := InterfaceOptions()

	serverConfig := config.Server
	ServerNodeTableInit(&serverConfig)

	FromNodeTableInit("", config.Clients)

	InterfaceGet := func() int {
		for i, addr := range interfaces {
			if addr == serverConfig.Endpoint {
				return i
			}
		}
		return 0
	}

	_, err := Dialog{
		AssignTo:      &dlg,
		Title:         "Server Configuration Dialog",
		Icon:          walk.IconInformation(),
		MinSize:       Size{Width: 800, Height: 500},
		Size:          Size{Width: 800, Height: 500},
		Font:          DefaultFont(),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Background: SolidColorBrush{Color: walk.RGB(220, 220, 220)},
				Layout:     HBox{},
				Children: []Widget{
					Label{
						Text: "Server Name:",
					},
					LineEdit{
						AssignTo: &nameLine,
						MinSize:  Size{Width: 100},
						Text:     serverConfig.Name,
						OnEditingFinished: func() {
							serverConfig.Name = nameLine.Text()
						},
					},
					Label{
						Text: "Server Address:",
					},
					ComboBox{
						AssignTo:     &listenBox,
						CurrentIndex: InterfaceGet(),
						Model:        interfaces,
						OnCurrentIndexChanged: func() {
							serverConfig.Endpoint = listenBox.Text()
						},
						OnBoundsChanged: func() {
							listenBox.SetCurrentIndex(InterfaceGet())
						},
					},
					Label{
						Text: "Server Port:",
					},
					NumberEdit{
						AssignTo:    &portNum,
						MinSize:     Size{Width: 100},
						Value:       float64(serverConfig.Port),
						ToolTipText: "0~65535",
						MaxValue:    65535,
						MinValue:    0,
						OnValueChanged: func() {
							serverConfig.Port = int(portNum.Value())
						},
					},
					CheckBox{
						AssignTo: &enableCB,
						Text:     "Enable",
						Checked:  serverConfig.Enable,
						OnCheckedChanged: func() {
							serverConfig.Enable = enableCB.Checked()
						},
					},
					HSpacer{
						MinSize: Size{Width: 50},
					},
					PushButton{
						AssignTo: &startTest,
						Text:     "Start Testing",
						OnClicked: func() {
							var err error
							server, err = ServerStartupTest(serverConfig)
							if err != nil {
								ErrorBoxAction(dlg, "Service startup failed! Reasons are as follows:"+err.Error())
								return
							}
							startTest.SetEnabled(false)
							stopTest.SetEnabled(true)
						},
					},
					PushButton{
						AssignTo: &stopTest,
						Text:     "Stop Testing",
						Enabled:  false,
						OnClicked: func() {
							if server != nil {
								server.Close()
								server = nil
							}
							startTest.SetEnabled(true)
							stopTest.SetEnabled(false)
						},
					},
				},
			},
			HSplitter{
				Children: []Widget{
					Composite{
						Layout:        VBox{},
						Background:    SolidColorBrush{Color: walk.RGB(220, 220, 220)},
						StretchFactor: 1,
						Children: []Widget{
							Label{
								Text: "Client Node:",
							},
							TableView{
								AssignTo:         &fromNodeTableView,
								AlternatingRowBG: true,
								ColumnsOrderable: true,
								CheckBoxes:       true,
								Columns: []TableViewColumn{
									{Title: "#", Width: 40},
									{Title: "Client Name", Width: 100},
									{Title: "Node Tag", Width: 300},
								},
								StyleCell: func(style *walk.CellStyle) {
									if style.Row()%2 == 0 {
										style.BackgroundColor = walk.RGB(248, 248, 255)
									} else {
										style.BackgroundColor = walk.RGB(220, 220, 220)
									}
								},
								Model: fromNodeTable,
								OnItemActivated: func() {
								},
							},
							Composite{
								Layout: HBox{MarginsZero: true},
								Children: []Widget{
									CheckBox{
										AssignTo: &fromCheckBox,
										Text:     "Select",
										OnCheckedChanged: func() {
											FromNodeTableSelect(fromCheckBox.Checked())
										},
									},
									HSpacer{},
									Label{
										Text: "Load Node:",
									},
									ComboBox{
										AssignTo:     &clientBox,
										CurrentIndex: 0,
										Model: func() []string {
											clientList := []string{"All Item"}
											clientList = append(clientList, config.ClientNames()...)
											return clientList
										}(),
										OnCurrentIndexChanged: func() {
											index := clientBox.CurrentIndex()
											if index == 0 {
												FromNodeTableInit("", config.Clients)
											} else {
												FromNodeTableInit(clientBox.Text(), config.Clients)
											}
										},
									},
									HSpacer{},
									PushButton{
										Text: "Add Select Nodes",
										OnClicked: func() {
											ServerNodeTableAdd(&serverConfig, fromNodeTable.SelectItems())
										},
									},
									PushButton{
										Text: "Add All Nodes",
										OnClicked: func() {
											ServerNodeTableAdd(&serverConfig, fromNodeTable.items)
										},
									},
								},
							},
						},
					},
					Composite{
						Layout:        VBox{},
						Background:    SolidColorBrush{Color: walk.RGB(220, 220, 220)},
						StretchFactor: 3,
						Children: []Widget{
							Label{
								Text: "Server Node:",
							},
							TableView{
								AssignTo:         &toNodeTableView,
								AlternatingRowBG: true,
								ColumnsOrderable: true,
								CheckBoxes:       true,
								Columns: []TableViewColumn{
									{Title: "#", Width: 40},
									{Title: "Client Name", Width: 120},
									{Title: "Client Node Tag", Width: 200},
									{Title: "Server Node Tag", Width: 200},
								},
								StyleCell: func(style *walk.CellStyle) {
									if style.Row()%2 == 0 {
										style.BackgroundColor = walk.RGB(248, 248, 255)
									} else {
										style.BackgroundColor = walk.RGB(220, 220, 220)
									}
								},
								Model: serverNodeTable,
								OnItemActivated: func() {
									index := toNodeTableView.CurrentIndex()
									if index < len(serverNodeTable.items) {
										ServerNodeEditDialog(dlg, &serverConfig, serverNodeTable.items[index])
									}
								},
							},
							Composite{
								Layout: HBox{MarginsZero: true},
								Children: []Widget{
									CheckBox{
										AssignTo: &serverCheckBox,
										Text:     "Select",
										OnCheckedChanged: func() {
											ServerNodeTableSelect(serverCheckBox.Checked())
										},
									},
									HSpacer{},
									PushButton{
										Text: "Delete Select",
										OnClicked: func() {
											ServerNodeTableDelete(&serverConfig)
										},
									},
									PushButton{
										Text: "All Select",
										OnClicked: func() {
											ServerNodeTableClean(&serverConfig)
										},
									},
								},
							},
						},
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "Accept",
						OnClicked: func() {
							if nameLine.Text() == "" {
								ErrorBoxAction(dlg, "Service name cannot be empty!")
								return
							}
							config.UpdateServer(serverConfig)
							dlg.Accept()
							logs.Info("server edit dialog accept")
						},
					},
					PushButton{
						AssignTo: &cancelPB,
						Text:     "Cancel",
						OnClicked: func() {
							dlg.Cancel()
							logs.Info("server edit dialog cancel")
						},
					},
				},
			},
		},
	}.Run(from)

	if err != nil {
		logs.Error("ServerEditDialog: %s", err.Error())
	}
}
