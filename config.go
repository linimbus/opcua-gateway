package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
)

type ApplicationLogConfig struct {
	Filename string `json:"filename"`
	Level    int    `json:"level"`
	MaxLines int    `json:"maxlines"`
	MaxSize  int    `json:"maxsize"`
	Daily    bool   `json:"daily"`
	MaxDays  int    `json:"maxdays"`
	Color    bool   `json:"color"`
}

type ApplicationConfig struct {
	Startup   bool                 `json:"startup"`
	LastPath  string               `json:"lastpath"`
	LogPath   string               `json:"logpath"`
	LogConfig ApplicationLogConfig `json:"logconfig"`
}

type DataStoreConfig struct {
	Enable   bool   `json:"enable"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	DataBase string `json:"database"`
	Expired  int    `json:"expired"`
}

type ClientConfig struct {
	Enable   bool       `json:"enable"`
	Timeout  int        `json:"timeout"`
	Name     string     `json:"name"`
	Endpoint string     `json:"endpoint"`
	Store    bool       `json:"store"`
	NodeList []NodeInfo `json:"nodes"`
}

type ServerNodeInfo struct {
	ClientName     string   `json:"clientName"`
	ClientEndpoint string   `json:"clientEndpoint"`
	ClientNode     NodeInfo `json:"clientNode"`
	ServerName     string   `json:"serverName"`
	ServerNode     NodeInfo `json:"serverNode"`
}

type ServerConfig struct {
	Enable   bool             `json:"enable"`
	Name     string           `json:"name"`
	Endpoint string           `json:"endpoint"`
	Port     int              `json:"port"`
	NodeList []ServerNodeInfo `json:"nodes"`
}

type Config struct {
	Filepath  string          `json:"-"`
	Startup   bool            `json:"startup"`
	Clients   []ClientConfig  `json:"clients"`
	Server    ServerConfig    `json:"server"`
	Datastore DataStoreConfig `json:"datastore"`
}

var defaultApplicationConfig = ApplicationConfig{
	Startup:  false,
	LastPath: "",
	LogPath:  "",
	LogConfig: ApplicationLogConfig{
		Filename: "runlog.log",
		Level:    logs.LevelInformational,
		Daily:    true,
		MaxSize:  100 * 1024 * 1024,
		MaxLines: 100 * 1024,
		MaxDays:  7,
		Color:    false,
	},
}

var defaultConfig = Config{
	Filepath: "",
	Startup:  false,
	Clients:  make([]ClientConfig, 0),
	Server:   ServerConfig{Enable: false, Name: "Server", Endpoint: "0.0.0.0", Port: 4840},
	Datastore: DataStoreConfig{
		Enable:  false,
		Address: "localhost", Port: 3306,
		UserName: "root", PassWord: "root",
		DataBase: "opcua", Expired: 30},
}

func (c *Config) statusUpdate() {
	body, err := os.ReadFile(c.Filepath)
	if err != nil {
		logs.Error("read config file from app data dir fail, %s", err.Error())
		return
	}

	value, err := json.Marshal(*c)
	if err != nil {
		logs.Error("covert config data to json fail, %s", err.Error())
		return
	}

	if string(body) == string(value) {
		StatusConfig(c.Filepath)
	} else {
		StatusConfig(c.Filepath + "*")
	}
}

func (c *Config) UpdateStartup(startup bool) {
	defer c.statusUpdate()
	c.Startup = startup
}

func (c *Config) UpdateDataBase(database DataStoreConfig) {
	defer c.statusUpdate()
	c.Datastore = database
}

func (c *Config) UpdateServer(server ServerConfig) {
	defer c.statusUpdate()
	c.Server = server
}

func (c *Config) UpdateClient(clients []ClientConfig) {
	defer c.statusUpdate()
	c.Clients = clients
}

func (c *Config) Add(client ClientConfig) error {
	defer c.statusUpdate()

	for _, item := range c.Clients {
		if item.Name == client.Name {
			return fmt.Errorf("client name %s already exist", client.Name)
		}
	}
	c.Clients = append(c.Clients, client)
	return nil
}

func (c *Config) Update(client ClientConfig) error {
	defer c.statusUpdate()

	for i, item := range c.Clients {
		if item.Name == client.Name {
			c.Clients[i] = client
			return nil
		}
	}
	return fmt.Errorf("client name %s not exist", client.Name)
}

func (c *Config) Delete(name string) {
	defer c.statusUpdate()

	for i, client := range c.Clients {
		if client.Name == name {
			c.Clients = append(c.Clients[:i], c.Clients[i+1:]...)
			return
		}
	}
}

func (c *Config) Save() error {
	defer c.statusUpdate()

	value, err := json.Marshal(*c)
	if err != nil {
		logs.Error("covert config data to json fail, %s", err.Error())
		return err
	}
	err = os.WriteFile(c.Filepath, value, 0664)
	if err != nil {
		logs.Error("write config data to file fail, %s", err.Error())
		return err
	}
	logs.Info("config save %s", string(value))

	return nil
}

func (c *Config) ClientConfig(name string) ClientConfig {
	for _, client := range c.Clients {
		if client.Name == name {
			return client
		}
	}
	return ClientConfig{}
}

func (c *Config) ClientNames(filters ...string) []string {
	names := make([]string, 0)
	for _, client := range c.Clients {
		if len(filters) == 0 {
			names = append(names, client.Name)
		} else {
			var flag bool
			for _, name := range filters {
				if client.Name == name {
					flag = true
				}
			}
			if !flag {
				names = append(names, client.Name)
			}
		}
	}
	return names
}

func (c *ServerConfig) Add(name string, endpoint string, node NodeInfo) bool {
	serverName := fmt.Sprintf("%s.%s", name, node.NodeID)
	for _, node := range c.NodeList {
		if node.ServerName == serverName {
			return false
		}
	}
	c.NodeList = append(c.NodeList, ServerNodeInfo{
		ClientName:     name,
		ClientEndpoint: endpoint,
		ClientNode:     node,
		ServerName:     serverName,
		ServerNode:     NodeInfo{NsIndex: node.NsIndex, NodeID: serverName},
	})
	return true
}

func (c *ServerConfig) Update(serverName string, node NodeInfo) bool {
	for i, item := range c.NodeList {
		if item.ServerName == serverName {
			c.NodeList[i].ServerNode = node
			return true
		}
	}
	return false
}

func (c *ServerConfig) Delete(serverName string) bool {
	for i, node := range c.NodeList {
		if node.ServerName == serverName {
			c.NodeList = append(c.NodeList[:i], c.NodeList[i+1:]...)
			return true
		}
	}
	return false
}

func (c *ServerConfig) Clean() {
	c.NodeList = make([]ServerNodeInfo, 0)
	logs.Info("server config node list clean")
}

func (c *ClientConfig) Reset(nodes []NodeInfo) {
	c.NodeList = nodes
}

func ConfigCreate(filepath string) (*Config, error) {
	config := defaultConfig
	config.Filepath = filepath
	err := config.Save()
	if err != nil {
		logs.Error("save default config to file fail, %s", err.Error())
		return nil, err
	}
	config.statusUpdate()
	return &config, nil
}

func ConfigLoad(filepath string) (*Config, error) {
	value, err := os.ReadFile(filepath)
	if err != nil {
		logs.Error("read config file from app data dir fail, %s", err.Error())
		return nil, err
	}
	logs.Info("config load %s value %s", filepath, string(value))

	var config Config
	err = json.Unmarshal(value, &config)
	if err != nil {
		logs.Error("json unmarshal config fail, %s", err.Error())
		return nil, err
	}
	config.Filepath = filepath
	config.statusUpdate()
	return &config, nil
}

func FileDialogOpen(from walk.Form, prevFilePath string) (string, error) {
	dlg := new(walk.FileDialog)

	dlg.FilePath = prevFilePath
	dlg.Filter = "*.json|*.json"
	dlg.Title = "Please select a configuration file"

	if ok, err := dlg.ShowOpen(from); err != nil {
		return "", err
	} else if !ok {
		return "", nil
	}

	logs.Info("config file dialog open %s", dlg.FilePath)

	return dlg.FilePath, nil
}

func FileDialogSave(from walk.Form, prevFilePath string) (string, error) {
	dlg := new(walk.FileDialog)

	dlg.FilePath = prevFilePath
	dlg.Filter = "*.json|*.json"
	dlg.Title = "Create an empty configuration file"

	if ok, err := dlg.ShowSave(from); err != nil {
		return "", err
	} else if !ok {
		return "", nil
	}
	logs.Info("config file dialog save %s.json", dlg.FilePath)

	return dlg.FilePath + ".json", nil
}

func (cfg *ApplicationConfig) Load() error {
	body, err := os.ReadFile(filepath.Join(ConfigDirGet(), "config.json"))
	if err != nil {
		fmt.Printf("read log config file from app data dir fail, %s", err.Error())
		return err
	}
	var config ApplicationConfig
	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Printf("json unmarshal log config fail, %s", err.Error())
		return err
	}
	*cfg = config
	return nil
}

func (cfg *ApplicationConfig) Save() {
	value, err := json.Marshal(*cfg)
	if err != nil {
		fmt.Printf("json marshal log config fail, %s", err.Error())
		return
	}
	err = os.WriteFile(filepath.Join(ConfigDirGet(), "config.json"), value, 0664)
	if err != nil {
		fmt.Printf("write log config to file fail, %s", err.Error())
		return
	}
}

func (cfg *ApplicationConfig) String() string {
	value, err := json.Marshal(*cfg)
	if err != nil {
		return ""
	}
	return string(value)
}

func (cfg *ApplicationConfig) UpdateStartup(startup bool) {
	cfg.Startup = startup
	cfg.Save()
}

func (cfg *ApplicationConfig) UpdateLastPath(path string) {
	cfg.LastPath = path
	cfg.Save()
}

func (cfg *ApplicationConfig) UpdateLogPath(path string) {
	cfg.LogPath = path
	cfg.Save()
	ApplicationLogInit(defaultApplicationConfig)
}

func ConfigInit() {
	err := defaultApplicationConfig.Load()
	if err != nil {
		defaultApplicationConfig.Save()
	}
	if defaultApplicationConfig.LogPath == "" {
		defaultApplicationConfig.LogPath = RunlogDirGet()
		defaultApplicationConfig.Save()
	}
	ApplicationLogInit(defaultApplicationConfig)
	logs.Info("config init %s", defaultApplicationConfig.String())
}
