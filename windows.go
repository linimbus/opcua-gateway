package main

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type StatItem struct {
	Name string

	Status   bool
	OperOK   uint64
	OperFail uint64

	checked bool
}

func (s *StatItem) Clear() {
	s.Status = false
	s.OperFail = 0
	s.OperOK = 0
}

type StatTable struct {
	sync.RWMutex

	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder

	items []*StatItem
}

func (n *StatTable) RowCount() int {
	return len(n.items)
}

func (n *StatTable) Value(row, col int) interface{} {
	item := n.items[row]
	switch col {
	case 0:
		return item.Name
	case 1:
		if item.Status {
			return "activate"
		}
		return "inactive"
	case 2:
		return item.OperOK
	case 3:
		return item.OperFail
	}
	panic("unexpected col")
}

func (n *StatTable) Checked(row int) bool {
	return n.items[row].checked
}

func (n *StatTable) SetChecked(row int, checked bool) error {
	n.items[row].checked = checked
	return nil
}

func (m *StatTable) Sort(col int, order walk.SortOrder) error {
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
			return c(a.Name < b.Name)
		case 1:
			return c(a.Status)
		case 2:
			return c(a.OperOK < b.OperOK)
		case 3:
			return c(a.OperFail < b.OperFail)
		}
		panic("unreachable")
	})
	return m.SorterBase.Sort(col, order)
}

const (
	STAT_CLIENT = "OPCUA Client"
	STAT_SERVER = "OPCUA Server"
	STAT_MYSQL  = "MYSQL Data Store"
)

var mainWindow *walk.MainWindow
var saveAction, clientEditAction, serverEditAction, mysqlEditAction *walk.Action
var statTableView *walk.TableView
var startPB, stopPB *walk.PushButton
var globalConfig *Config
var timestampView *walk.Label
var startupBox *walk.CheckBox
var globalStat *StatTable
var instance *OpcuaServer

func ConfigLoadAuto() {
	for {
		if ActionReady() {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Millisecond * 500)

	if defaultApplicationConfig.LastPath != "" {
		logs.Info("auto load config file %s", defaultApplicationConfig.LastPath)
		config, err := ConfigLoad(defaultApplicationConfig.LastPath)
		if err != nil {
			logs.Warning("auto load config file fail, %s", err.Error())
			return
		}
		globalConfig = config
		logs.Info("auto load config file success")

		if defaultApplicationConfig.Startup && !ServerRunning() {
			logs.Info("server auto startup")
			ServerSwitch()
		}
		ActionEnable()
	}
}

func init() {
	globalStat = new(StatTable)
	globalStat.items = make([]*StatItem, 0)
	globalStat.items = append(globalStat.items, &StatItem{Name: STAT_CLIENT})
	globalStat.items = append(globalStat.items, &StatItem{Name: STAT_SERVER})
	globalStat.items = append(globalStat.items, &StatItem{Name: STAT_MYSQL})
}

func StatUpdateTask() {
	logs.Info("stat update task running")
	for {
		time.Sleep(time.Millisecond * 500)

		globalStat.Lock()
		globalStat.PublishRowsReset()
		globalStat.Sort(globalStat.sortColumn, globalStat.sortOrder)
		globalStat.Unlock()

		if timestampView != nil && timestampView.Visible() {
			timestampView.SetText(TimeStampGet())
		}
	}
}

func StatClear() {
	globalStat.Lock()
	defer globalStat.Unlock()

	for _, item := range globalStat.items {
		atomic.StoreUint64(&item.OperFail, 0)
		atomic.StoreUint64(&item.OperOK, 0)
	}
}

func ActionReady() bool {
	if saveAction != nil &&
		clientEditAction != nil &&
		serverEditAction != nil &&
		mysqlEditAction != nil {
		return true
	}
	return false
}

func ActionEnable() {
	saveAction.SetEnabled(true)
	clientEditAction.SetEnabled(true)
	serverEditAction.SetEnabled(true)
	mysqlEditAction.SetEnabled(true)
}

func MenuBarInit() []MenuItem {
	return []MenuItem{
		Menu{
			Text: "Configuration Files",
			Items: []MenuItem{
				Action{
					Text:     "Create configuration file",
					Shortcut: Shortcut{Modifiers: walk.ModControl, Key: walk.KeyN},

					OnTriggered: func() {
						fileName, err := FileDialogSave(mainWindow, "")
						if err != nil {
							ErrorBoxAction(mainWindow, err.Error())
							return
						}
						if fileName != "" {
							logs.Info("create config file %s", fileName)

							config, err := ConfigCreate(fileName)
							if err != nil {
								ErrorBoxAction(mainWindow, "Configuration file creation fails for the following reason:"+err.Error())
								return
							}
							globalConfig = config
							ActionEnable()
						}
					},
				},
				Action{
					Text:     "Open configuration file",
					Shortcut: Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
					OnTriggered: func() {
						fileName, err := FileDialogOpen(mainWindow, "")
						if err != nil {
							ErrorBoxAction(mainWindow, err.Error())
							return
						}

						if fileName != "" {
							logs.Info("ready to open config file %s", fileName)

							config, err := ConfigLoad(fileName)
							if err != nil {
								ErrorBoxAction(mainWindow, "Configuration file loading fails for the following reason:"+err.Error())
								return
							}
							globalConfig = config
							defaultApplicationConfig.UpdateLastPath(globalConfig.Filepath)

							ActionEnable()
						}
					},
				},
				Action{
					Text:     "Save Configuration File",
					AssignTo: &saveAction,
					Enabled:  false,
					Shortcut: Shortcut{Modifiers: walk.ModControl, Key: walk.KeyS},
					OnTriggered: func() {
						err := globalConfig.Save()
						if err != nil {
							ErrorBoxAction(mainWindow, "Configuration file saving fails for the following reason:"+err.Error())
						} else {
							InfoBoxAction(mainWindow, "Configuration file saved successfully")

							defaultApplicationConfig.UpdateLastPath(globalConfig.Filepath)
						}
					},
				},
				Separator{},
				Action{
					Text: "Exit",
					OnTriggered: func() {
						ConfirmBoxAction(mainWindow, "Shutdown the program?", CloseWindows)
					},
				},
			},
		},
		Menu{
			Text: "Configuration Editor",
			Items: []MenuItem{
				Action{
					AssignTo: &clientEditAction,
					Text:     "OPCUA Client Settings",
					Enabled:  false,
					OnTriggered: func() {
						ClientDialog(mainWindow, globalConfig)
					},
				},
				Action{
					AssignTo: &serverEditAction,
					Text:     "OPCUA Server Settings",
					Enabled:  false,
					OnTriggered: func() {
						ServerEditDialog(mainWindow, globalConfig)
					},
				},
				Action{
					AssignTo: &mysqlEditAction,
					Text:     "MYSQL Database Settings",
					Enabled:  false,
					OnTriggered: func() {
						DataStoreDialog(mainWindow, globalConfig)
					},
				},
			},
		},
		Action{
			Text: "Diagnostic Log",
			OnTriggered: func() {
				OpenBrowserWeb(RunlogDirGet())
			},
		},
		Action{
			Text: "Hide Windows",
			OnTriggered: func() {
				NotifyAction()
			},
		},
		Action{
			Text: "About",
			OnTriggered: func() {
				AboutAction()
			},
		},
	}
}

func ServerRunning() bool {
	return instance != nil
}

func ServerStart() error {
	if globalConfig == nil {
		return fmt.Errorf("Configuration file not loaded")
	}

	if len(globalConfig.Clients) == 0 {
		return fmt.Errorf("No OPCUA client is configured")
	}

	var err error
	instance, err = NewOpcuaServer(*globalConfig, globalStat.items)
	if err != nil {
		logs.Error("server start failed, %s", err.Error())
		return fmt.Errorf("Starting the service fails for the following reason: %s", err.Error())
	}

	return nil
}

func ServerShutdown() error {
	instance.Close()
	instance = nil
	return nil
}

func ServerStatus(flag bool) {
	startPB.SetEnabled(!flag)
	stopPB.SetEnabled(flag)
}

func ServerSwitch() {
	var err error

	startPB.SetEnabled(false)
	stopPB.SetEnabled(false)

	time.Sleep(time.Millisecond * 200)

	if ServerRunning() {
		err = ServerShutdown()
	} else {
		err = ServerStart()
	}
	if err != nil {
		ErrorBoxAction(mainWindow, err.Error())
	}

	ServerStatus(ServerRunning())
}

func ConsoleWidget() []Widget {
	go StatUpdateTask()

	return []Widget{
		Composite{
			Layout: HBox{MarginsZero: true},
			Children: []Widget{
				Label{
					AssignTo:    &timestampView,
					ToolTipText: "Current system time",
					Font:        Font{Family: "Segoe UI", Bold: true, PointSize: 14},
					Text:        TimeStampGet(),
				},
				HSpacer{},
				CheckBox{
					AssignTo:    &startupBox,
					ToolTipText: "Next automatic service startup",
					Text:        "Auto Startup",
					Checked:     defaultApplicationConfig.Startup,
					OnCheckedChanged: func() {
						defaultApplicationConfig.UpdateStartup(startupBox.Checked())
					},
				},
				PushButton{
					ToolTipText: "Service Statistics Count Zero",
					Text:        "Clear Count",
					OnClicked: func() {
						StatClear()
					},
				},
			},
		},
		TableView{
			AssignTo:                 &statTableView,
			ToolTipText:              "Statistical counts of the operation of services",
			AlternatingRowBG:         false,
			ColumnsOrderable:         false,
			CheckBoxes:               false,
			NotSortableByHeaderClick: true,
			Font:                     Font{Family: "Segoe UI", Bold: false, PointSize: 10},
			Columns: []TableViewColumn{
				{Title: "#", Width: 120},
				{Title: "Service Status", Width: 120},
				{Title: "Operation Success Count", Width: 200},
				{Title: "Operation Failure Count", Width: 200},
			},
			Model: globalStat,
		},
		Composite{
			Layout: HBox{MarginsZero: true},
			Children: []Widget{
				PushButton{
					AssignTo:    &startPB,
					Image:       ICON_Start,
					Text:        " ",
					Enabled:     true,
					ToolTipText: "Startup all services",
					MinSize:     Size{Height: 64},
					OnClicked: func() {
						ServerSwitch()
					},
				},
				PushButton{
					AssignTo:    &stopPB,
					Image:       ICON_Stop,
					Text:        " ",
					Enabled:     false,
					ToolTipText: "Stop all services",
					MinSize:     Size{Height: 64},
					OnClicked: func() {
						ServerSwitch()
					},
				},
			},
		},
	}
}

func CreateWindow() {
	CapSignal(CloseWindows)

	logs.Info("create windows success")

	go ConfigLoadAuto()

	cnt, err := MainWindow{
		Title:          AppNameGet(),
		Icon:           ICON_Main,
		AssignTo:       &mainWindow,
		MinSize:        Size{Width: 600, Height: 300},
		Size:           Size{Width: 650, Height: 350},
		Layout:         VBox{},
		Font:           DefaultFont(),
		MenuItems:      MenuBarInit(),
		StatusBarItems: StatusBarInit(),
		Children: []Widget{
			Composite{
				Layout:   VBox{MarginsZero: true},
				Children: ConsoleWidget(),
			},
		},
	}.Run()

	if err != nil {
		logs.Error("main windows exit %s", err.Error())
	} else {
		logs.Info("main windows exit %d", cnt)
	}

	if err := recover(); err != nil {
		logs.Error("main windows panic %v", err)
	}

	logs.Info("close windows")

	CloseWindows()
}

func CloseWindows() {
	if ServerRunning() {
		ServerShutdown()
	}
	if mainWindow != nil {
		mainWindow.Close()
		mainWindow = nil
	}
	NotifyExit()
}
