package main

import (
	"fmt"
	"sort"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type ClientItem struct {
	Index int

	Client  ClientConfig
	checked bool
}

type ClientTable struct {
	sync.RWMutex

	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder

	items []*ClientItem
}

func (n *ClientTable) RowCount() int {
	return len(n.items)
}

func (n *ClientTable) Value(row, col int) interface{} {
	item := n.items[row]
	switch col {
	case 0:
		return item.Index
	case 1:
		return item.Client.Name
	case 2:
		return item.Client.Endpoint
	case 3:
		return fmt.Sprintf("%d ms", item.Client.Timeout)
	case 4:
		return len(item.Client.NodeList)
	case 5:
		return SwitchName(item.Client.Enable)
	case 6:
		return SwitchName(item.Client.Store)
	}
	panic("unexpected col")
}

func (n *ClientTable) Checked(row int) bool {
	return n.items[row].checked
}

func (n *ClientTable) SetChecked(row int, checked bool) error {
	n.items[row].checked = checked
	return nil
}

func (m *ClientTable) Sort(col int, order walk.SortOrder) error {
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
			return c(a.Client.Name < b.Client.Name)
		case 2:
			return c(a.Client.Endpoint < b.Client.Endpoint)
		case 3:
			return c(a.Client.Timeout < b.Client.Timeout)
		case 4:
			return c(len(a.Client.NodeList) < len(b.Client.NodeList))
		case 5:
			return c(a.Client.Enable)
		case 6:
			return c(a.Client.Store)
		}
		panic("unreachable")
	})
	return m.SorterBase.Sort(col, order)
}

var clientTable *ClientTable
var clientTableView *walk.TableView

func init() {
	clientTable = new(ClientTable)
	clientTable.items = make([]*ClientItem, 0)
}

func ClientTableInit(clients []ClientConfig) {
	clientTable.Lock()
	defer clientTable.Unlock()

	item := make([]*ClientItem, 0)
	for i, client := range clients {
		item = append(item, &ClientItem{Index: i, Client: client})
	}
	clientTable.items = item
	clientTable.PublishRowsReset()
	clientTable.Sort(clientTable.sortColumn, clientTable.sortOrder)
}

func ClientTableDelete(config *Config) {
	clientTable.Lock()
	defer clientTable.Unlock()

	clients := make([]ClientConfig, 0)
	item := make([]*ClientItem, 0)
	for _, v := range clientTable.items {
		if !v.checked {
			item = append(item, v)
			clients = append(clients, v.Client)
		}
	}

	config.UpdateClient(clients)

	clientTable.items = item
	clientTable.PublishRowsReset()
	clientTable.Sort(clientTable.sortColumn, clientTable.sortOrder)
}

func ClientTableSelect() *ClientItem {
	clientTable.Lock()
	defer clientTable.Unlock()

	for _, v := range clientTable.items {
		if v.checked {
			return v
		}
	}
	return nil
}

func ClientAddDialog(from walk.Form, config *Config) {
	var dlg *walk.Dialog
	var nameLine, addressLine *walk.LineEdit
	var enableCB, storeCB *walk.CheckBox
	var acceptPB, cancelPB, testPB *walk.PushButton
	var number *walk.NumberEdit

	_, err := Dialog{
		AssignTo:      &dlg,
		Title:         "Adding a Client Configuration",
		Icon:          walk.IconInformation(),
		MinSize:       Size{Width: 600, Height: 300},
		Size:          Size{Width: 600, Height: 300},
		Font:          DefaultFont(),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "Client Name:",
					},
					LineEdit{
						AssignTo: &nameLine,
					},
					Label{
						Text: "Data Collection Address:",
					},
					LineEdit{
						AssignTo: &addressLine,
					},
					Label{
						Text: "Data Collection Frequency:",
					},
					NumberEdit{
						AssignTo:    &number,
						Value:       float64(1000),
						ToolTipText: "100~10000 ms",
						MaxValue:    10000,
						MinValue:    100,
					},
					HSpacer{},
					Composite{
						Layout: HBox{},
						Children: []Widget{
							CheckBox{
								Text:     "Collection Enable",
								AssignTo: &enableCB,
								Checked:  true,
							},
							CheckBox{
								Text:     "Data Store",
								AssignTo: &storeCB,
								Checked:  true,
							},
							PushButton{
								AssignTo: &testPB,
								Text:     "Connectivity Test",
								OnClicked: func() {
									testPB.SetEnabled(false)
									defer testPB.SetEnabled(true)

									err := ClientTest(addressLine.Text())
									if err != nil {
										ErrorBoxAction(dlg, "OPCUA connection failed:"+err.Error())
									} else {
										InfoBoxAction(dlg, "OPC Connection Successful!")
									}
								},
							},
						},
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
							if nameLine.Text() == "" {
								ErrorBoxAction(dlg, "The client name cannot be empty!")
								return
							}
							if addressLine.Text() == "" {
								ErrorBoxAction(dlg, "The data collection address cannot be empty!")
								return
							}

							err := config.Add(ClientConfig{
								Name:     nameLine.Text(),
								Endpoint: addressLine.Text(),
								Timeout:  int(number.Value()),
								Enable:   enableCB.Checked(),
								Store:    storeCB.Checked()})

							if err != nil {
								ErrorBoxAction(dlg, "Add client failed: "+err.Error())
								return
							}
							ClientTableInit(config.Clients)
							dlg.Accept()
							logs.Info("client add dialog accept")
						},
					},
					PushButton{
						AssignTo: &cancelPB,
						Text:     "Cancel",
						OnClicked: func() {
							dlg.Cancel()
							logs.Info("client add dialog cancel")
						},
					},
					HSpacer{},
				},
			},
		},
	}.Run(from)

	if err != nil {
		logs.Error("ClientAddDialog: %s", err.Error())
	}
}

func ClientDialog(from walk.Form, config *Config) {
	var dlg *walk.Dialog
	var addPB, editPB, deletePB *walk.PushButton

	ClientTableInit(config.Clients)

	_, err := Dialog{
		AssignTo: &dlg,
		Title:    "OPCUA Client Configuration",
		Icon:     walk.IconInformation(),
		MinSize:  Size{Width: 800, Height: 400},
		Size:     Size{Width: 800, Height: 400},
		Font:     DefaultFont(),
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text: "Client List:",
			},
			TableView{
				AssignTo:         &clientTableView,
				AlternatingRowBG: true,
				ColumnsOrderable: true,
				CheckBoxes:       true,
				Columns: []TableViewColumn{
					{Title: "#", Width: 60},
					{Title: "Client Name", Width: 120},
					{Title: "Data Collection address", Width: 200},
					{Title: "Data Collection Frequency", Width: 100},
					{Title: "Number Nodes", Width: 80},
					{Title: "Collection Enable", Width: 80},
					{Title: "Data Store", Width: 80},
				},
				StyleCell: func(style *walk.CellStyle) {
					if style.Row()%2 == 0 {
						style.BackgroundColor = walk.RGB(248, 248, 255)
					} else {
						style.BackgroundColor = walk.RGB(220, 220, 220)
					}
				},
				Model: clientTable,
				OnItemActivated: func() {
					index := clientTableView.CurrentIndex()
					if index >= len(clientTable.items) {
						ErrorBoxAction(from, "The edited client does not exist!")
						return
					}
					ClientNodeEditDialog(dlg, clientTable.items[index], config)
				},
			},
			VSpacer{},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo: &addPB,
						Text:     "Add",
						OnClicked: func() {
							ClientAddDialog(dlg, config)
						},
					},
					PushButton{
						AssignTo: &editPB,
						Text:     "Edit",
						OnClicked: func() {
							item := ClientTableSelect()
							if item == nil {
								InfoBoxAction(dlg, "Please select an item")
							} else {
								ClientNodeEditDialog(dlg, item, config)
							}
						},
					},
					PushButton{
						AssignTo: &deletePB,
						Text:     "Delete",
						OnClicked: func() {
							ClientTableDelete(config)
						},
					},
					HSpacer{},
				},
			},
		},
	}.Run(from)

	if err != nil {
		logs.Error("ClientDialog: %s", err.Error())
	}
}
