package main

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type NodeItem struct {
	Index int

	node  NodeInfo
	value string

	checked bool
}

type NodeTable struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder

	items []*NodeItem
}

func (n *NodeTable) RowCount() int {
	return len(n.items)
}

func (n *NodeTable) Value(row, col int) interface{} {
	item := n.items[row]
	switch col {
	case 0:
		return item.Index
	case 1:
		return item.node.ToString()
	case 2:
		return item.value
	}
	panic("unexpected col")
}

func (n *NodeTable) Checked(row int) bool {
	return n.items[row].checked
}

func (n *NodeTable) SetChecked(row int, checked bool) error {
	n.items[row].checked = checked
	return nil
}

func (m *NodeTable) Sort(col int, order walk.SortOrder) error {
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
			return c(a.node.ToString() < b.node.ToString())
		case 2:
			return c(a.value < b.value)
		}
		panic("unreachable")
	})
	return m.SorterBase.Sort(col, order)
}

func (m *NodeTable) Review() {
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
}

func (m *NodeTable) Save(config *ClientConfig) {
	nodes := make([]NodeInfo, 0)
	for _, item := range m.items {
		nodes = append(nodes, item.node)
	}
	config.Reset(nodes)
}

func (m *NodeTable) Init(config ClientConfig) {
	defer m.Review()

	items := make([]*NodeItem, 0)
	for index, node := range config.NodeList {
		items = append(items, &NodeItem{
			Index: index,
			node:  node,
			value: "",
		})
	}
	m.items = items
}

func (m *NodeTable) Delete(config *ClientConfig) {
	defer m.Save(config)
	defer m.Review()

	items := make([]*NodeItem, 0)
	for _, item := range m.items {
		if !item.checked {
			items = append(items, item)
		}
	}
	m.items = items
}

func (m *NodeTable) DeleteAll(config *ClientConfig) {
	defer m.Save(config)
	defer m.Review()

	m.items = make([]*NodeItem, 0)
}

func (m *NodeTable) Select(check bool) {
	defer m.Review()

	for _, node := range m.items {
		node.checked = check
	}
}

func (m *NodeTable) Order(index int, order OrderType, config *ClientConfig) {
	items := m.items
	if index < 0 || index >= len(items) {
		return
	}

	switch order {
	case ORDER_UP:
		if index > 0 {
			items[index-1], items[index] = items[index], items[index-1]
		} else {
			return
		}
	case ORDER_DOWN:
		if index < len(items)-1 {
			items[index], items[index+1] = items[index+1], items[index]
		} else {
			return
		}
	case ORDER_TOP:
		if index > 0 {
			temp := make([]*NodeItem, 0)
			temp = append(temp, items[index])
			temp = append(temp, items[:index]...)
			temp = append(temp, items[index+1:]...)
			items = temp
		} else {
			return
		}
	case ORDER_BOTTOM:
		if index < len(items)-1 {
			temp := make([]*NodeItem, 0)
			temp = append(temp, items[:index]...)
			temp = append(temp, items[index+1:]...)
			temp = append(temp, items[index])
			items = temp
		} else {
			return
		}
	}

	for i := 0; i < len(items); i++ {
		items[i].Index = i
	}

	m.items = items
	m.Save(config)
	m.Review()
}

func (m *NodeTable) ReadValue(config ClientConfig) error {
	client, err := NewClient(config.Endpoint)
	if err != nil {
		logs.Error("node table value read failed, %s", err.Error())
		return err
	}
	defer client.Close()

	for _, item := range m.items {
		node := item.node
		value, err := client.ReadNode(node)
		if err != nil {
			logs.Error("node table %d:%s value read failed, %s", node.ToString(), err.Error())
			continue
		}
		item.value = value.ToString()
	}
	m.Review()

	return nil
}

func (m *NodeTable) Query(node NodeInfo) bool {
	for _, item := range m.items {
		if item.node.NodeID == node.NodeID && item.node.NsIndex == node.NsIndex {
			return true
		}
	}
	return false
}

func (m *NodeTable) Append(nodes []NodeInfo, values []string, config *ClientConfig) {
	defer m.Save(config)
	defer m.Review()

	for i, node := range nodes {
		if m.Query(node) {
			continue
		}
		m.items = append(m.items, &NodeItem{
			Index: len(m.items),
			node:  NodeInfo{NsIndex: node.NsIndex, NodeID: node.NodeID},
			value: values[i],
		})
	}
}

type NodeTreeItem struct {
	nsIndex uint32
	nodeID  string

	parent   *NodeTreeItem
	children []*NodeTreeItem
}

var _ walk.TreeItem = new(NodeTreeItem)

func (d *NodeTreeItem) Text() string {
	return d.nodeID
}

func (d *NodeTreeItem) Parent() walk.TreeItem {
	if d.parent == nil {
		return nil
	}
	return d.parent
}

func (d *NodeTreeItem) ChildCount() int {
	return len(d.children)
}

func (d *NodeTreeItem) ChildAt(index int) walk.TreeItem {
	return d.children[index]
}

func (d *NodeTreeItem) Image() interface{} {
	return d.Path()
}

func (d *NodeTreeItem) ResetChildren() error {
	return nil
}

func (d *NodeTreeItem) Path() string {
	elems := []string{d.nodeID}

	dir, _ := d.Parent().(*NodeTreeItem)

	for dir != nil {
		elems = append([]string{dir.nodeID}, elems...)
		dir, _ = dir.Parent().(*NodeTreeItem)
	}

	return filepath.Join(elems...)
}

func (d *NodeTreeItem) Export(filter string) []*NodeTreeItem {
	output := make([]*NodeTreeItem, 0)

	if strings.Contains(d.nodeID, filter) {
		output = append(output, d)
	}

	for _, child := range d.children {
		output = append(output, child.Export(filter)...)
	}

	return output
}

type NodeTreeModel struct {
	walk.TreeModelBase
	roots []*NodeTreeItem
}

var _ walk.TreeModel = new(NodeTreeModel)

func (*NodeTreeModel) LazyPopulation() bool {
	return true
}

func (m *NodeTreeModel) RootCount() int {
	return len(m.roots)
}

func (m *NodeTreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}

var nodeTableView *walk.TableView
var nodeTreeView *walk.TreeView
var currentNodeTreeItem *NodeTreeItem

func NodeTreeItemInit(node *NodeTree, parent *NodeTreeItem, filter string, level, levelLimit int) *NodeTreeItem {
	item := &NodeTreeItem{
		nsIndex:  node.Node.NsIndex,
		nodeID:   node.Node.NodeID,
		parent:   parent,
		children: make([]*NodeTreeItem, 0),
	}

	if level >= levelLimit {
		logs.Info("node tree init search level %d >= %d", level, levelLimit)
		return item
	}

	for _, child := range node.SubNodes {
		item.children = append(item.children, NodeTreeItemInit(child, item, filter, level+1, levelLimit))
	}
	return item
}

func NodeTreeInit(nodes []*NodeTree, filter string, levelLimit int) {
	nodeTree := new(NodeTreeModel)
	nodeTree.roots = make([]*NodeTreeItem, 0)
	for _, node := range nodes {
		nodeTree.roots = append(nodeTree.roots, NodeTreeItemInit(node, nil, filter, 1, levelLimit))
	}
	nodeTreeView.SetModel(nodeTree)
	currentNodeTreeItem = nil
}

func NodeAddBatch(nodeTable *NodeTable, config *ClientConfig, roots []*NodeTreeItem, filter string) error {
	client, err := NewClient(config.Endpoint)
	if err != nil {
		logs.Error("export node tree for %s create client failed, %s", config.Name, err.Error())
		return err
	}
	defer client.Close()

	nodesExport := make([]*NodeTreeItem, 0)
	for _, root := range roots {
		nodesExport = append(nodesExport, root.Export(filter)...)
	}

	nodes := make([]NodeInfo, 0)
	values := make([]string, 0)

	for _, node := range nodesExport {
		nodeInfo := NodeInfo{NsIndex: node.nsIndex, NodeID: node.nodeID}
		value, err := client.ReadNode(nodeInfo)
		if err != nil {
			logs.Error("export node tree for %s read node %s failed, %s", config.Name, nodeInfo.ToString(), err.Error())
			continue
		}
		nodes = append(nodes, nodeInfo)
		values = append(values, value.ToString())

		logs.Info("export node tree for %s read node %s value %s success")
	}

	nodeTable.Append(nodes, values, config)
	return nil
}

func ClientNodeEditDialog(from walk.Form, clientItem *ClientItem, config *Config) {
	var dlg *walk.Dialog
	var acceptPB, cancelPB *walk.PushButton
	var endpoint, filterKey *walk.LineEdit
	var loadTreePB, addNodePB, addAllNodePB *walk.PushButton
	var deleteAllPB, deletePB, readValuePB *walk.PushButton
	var timeout, levelNumber *walk.NumberEdit
	var enable, store, selectBox *walk.CheckBox
	var nodeTable NodeTable

	client := clientItem.Client

	nodeTable.Init(client)

	nodeTree := new(NodeTreeModel)
	nodeTree.roots = make([]*NodeTreeItem, 0)

	_, err := Dialog{
		AssignTo:      &dlg,
		Title:         "Client Node Edit Dialog",
		Icon:          walk.IconInformation(),
		MinSize:       Size{Width: 1200, Height: 600},
		Size:          Size{Width: 1200, Height: 600},
		Font:          DefaultFont(),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout:     Grid{Columns: 4},
				Background: SolidColorBrush{Color: walk.RGB(220, 220, 220)},
				Children: []Widget{
					// line: 1
					Label{
						Text: "Client Name:",
					},
					LineEdit{
						Text:    client.Name,
						Enabled: false,
					},
					Label{
						Text: "Data collection frequency:",
					},
					NumberEdit{
						AssignTo:    &timeout,
						Value:       float64(client.Timeout),
						ToolTipText: "1~10000 ms",
						MinSize:     Size{Width: 60},
						MaxValue:    10000,
						MinValue:    1,
						OnValueChanged: func() {
							client.Timeout = int(timeout.Value())
						},
					},
					// line: 2
					Label{
						Text: "Data collection address:",
					},
					LineEdit{
						AssignTo: &endpoint,
						Text:     client.Endpoint,
						OnEditingFinished: func() {
							client.Endpoint = endpoint.Text()
						},
					},
					CheckBox{
						Text:     "Enable",
						AssignTo: &enable,
						Checked:  client.Enable,
						OnCheckedChanged: func() {
							client.Enable = enable.Checked()
						},
					},
					CheckBox{
						Text:     "Store",
						AssignTo: &store,
						Checked:  client.Store,
						OnCheckedChanged: func() {
							client.Store = store.Checked()
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
								Text: "Query Node Tree",
							},
							Composite{
								Layout: HBox{MarginsZero: true},
								Children: []Widget{
									Label{
										Text: "Traversal Level:",
									},
									NumberEdit{
										AssignTo:    &levelNumber,
										Value:       float64(7),
										ToolTipText: "1~100",
										MaxValue:    100,
										MinValue:    1,
										OnValueChanged: func() {
										},
									},
									PushButton{
										AssignTo: &loadTreePB,
										Text:     "Load Node",
										OnClicked: func() {
											go func() {
												loadTreePB.SetEnabled(false)
												defer loadTreePB.SetEnabled(true)

												cli, err := NewClient(endpoint.Text())
												if err != nil {
													ErrorBoxAction(dlg, "OPCUA connection failed:"+err.Error())
													return
												}
												defer cli.Close()

												nodeTree, err := cli.BrowseNode()
												if err != nil {
													ErrorBoxAction(dlg, "Failed to load node:"+err.Error())
													return
												}

												NodeTreeInit(nodeTree, filterKey.Text(), int(levelNumber.Value()))
											}()
										},
									},
								},
							},
							TreeView{
								AssignTo: &nodeTreeView,
								Model:    nodeTree,
								OnCurrentItemChanged: func() {
									currentNodeTreeItem = nodeTreeView.CurrentItem().(*NodeTreeItem)
								},
							},
							Composite{
								Layout: Grid{Columns: 2, MarginsZero: true},
								Children: []Widget{
									PushButton{
										AssignTo: &addNodePB,
										Text:     "Add Selected Node",
										OnClicked: func() {
											if currentNodeTreeItem == nil {
												ErrorBoxAction(dlg, "No nodes selected")
												return
											}

											err := NodeAddBatch(&nodeTable, &client, []*NodeTreeItem{currentNodeTreeItem}, "")
											if err != nil {
												ErrorBoxAction(dlg, "Error adding node, the reason is as follows:"+err.Error())
											}
										},
									},
									PushButton{
										AssignTo: &addAllNodePB,
										Text:     "Add All Node",
										OnClicked: func() {
											nodeTree := nodeTreeView.Model().(*NodeTreeModel)
											if len(nodeTree.roots) == 0 {
												ErrorBoxAction(dlg, "No nodes loaded")
												return
											}

											err := NodeAddBatch(&nodeTable, &client, nodeTree.roots, "")
											if err != nil {
												ErrorBoxAction(dlg, "Error adding node, the reason is as follows:"+err.Error())
											}
										},
									},
									Composite{
										Layout: HBox{MarginsZero: true},
										Children: []Widget{
											Label{
												Text: "Filter keyword:",
											},
											LineEdit{
												AssignTo: &filterKey,
												Text:     "",
												OnEditingFinished: func() {

												},
											},
										},
									},
									PushButton{
										AssignTo: &addAllNodePB,
										Text:     "Add Filter Node:",
										OnClicked: func() {
											nodeTree := nodeTreeView.Model().(*NodeTreeModel)
											if len(nodeTree.roots) == 0 {
												ErrorBoxAction(dlg, "No nodes loaded")
												return
											}
											err := NodeAddBatch(&nodeTable, &client, nodeTree.roots, filterKey.Text())
											if err != nil {
												ErrorBoxAction(dlg, "Error adding node, the reason is as follows:"+err.Error())
											}
										},
									},
								},
							},
						},
					},
					Composite{
						Layout:        VBox{},
						Background:    SolidColorBrush{Color: walk.RGB(220, 220, 220)},
						StretchFactor: 2,
						Children: []Widget{
							Label{
								Text: "Subscribe Node List:",
							},
							TableView{
								AssignTo:         &nodeTableView,
								AlternatingRowBG: true,
								ColumnsOrderable: true,
								CheckBoxes:       true,
								Columns: []TableViewColumn{
									{Title: "#", Width: 60},
									{Title: "Node Tag", Width: 300},
									{Title: "Node Data", Width: 200},
								},
								StyleCell: func(style *walk.CellStyle) {
									if style.Row()%2 == 0 {
										style.BackgroundColor = walk.RGB(248, 248, 255)
									} else {
										style.BackgroundColor = walk.RGB(220, 220, 220)
									}
								},
								Model: &nodeTable,
								OnItemActivated: func() {
								},
							},

							Composite{
								Layout: HBox{MarginsZero: true},
								Children: []Widget{
									CheckBox{
										AssignTo: &selectBox,
										Text:     "Select",
										OnClicked: func() {
											nodeTable.Select(selectBox.Checked())
										},
									},
									HSpacer{},
									PushButton{
										AssignTo: &deleteAllPB,
										Text:     "All Delete",
										OnClicked: func() {
											nodeTable.DeleteAll(&client)
										},
									},
									PushButton{
										AssignTo: &deletePB,
										Text:     "Selected Delete",
										OnClicked: func() {
											nodeTable.Delete(&client)
										},
									},
									HSpacer{},
									PushButton{
										Text: "Up",
										OnClicked: func() {
											nodeTable.Order(nodeTableView.CurrentIndex(), ORDER_UP, &client)
										},
									},
									PushButton{
										Text: "Down",
										OnClicked: func() {
											nodeTable.Order(nodeTableView.CurrentIndex(), ORDER_DOWN, &client)
										},
									},
									PushButton{
										Text: "Top",
										OnClicked: func() {
											nodeTable.Order(nodeTableView.CurrentIndex(), ORDER_TOP, &client)
										},
									},
									PushButton{
										Text: "Bottom",
										OnClicked: func() {
											nodeTable.Order(nodeTableView.CurrentIndex(), ORDER_BOTTOM, &client)
										},
									},
									HSpacer{},

									PushButton{
										AssignTo: &readValuePB,
										Text:     "Read Datas",
										OnClicked: func() {
											go func() {
												readValuePB.SetEnabled(false)
												defer readValuePB.SetEnabled(true)

												err := nodeTable.ReadValue(client)
												if err != nil {
													ErrorBoxAction(dlg, "Failed to read the data, following reasons:"+err.Error())
												}
											}()
										},
									},
									HSpacer{},
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
						Text: "Accept",
						OnClicked: func() {
							if client.Endpoint == "" {
								ErrorBoxAction(dlg, "The data collection address is empty")
								return
							}
							clientItem.Client = client
							config.Update(client)

							dlg.Accept()
							logs.Info("client node edit dialog accept")
						},
					},
					PushButton{
						Text: "Cancel",
						OnClicked: func() {
							dlg.Cancel()
							logs.Info("client node edit dialog cancel")
						},
					},
				},
			},
		},
	}.Run(from)

	if err != nil {
		logs.Error("ClientNodeEditDialog: %s", err.Error())
	}
}
