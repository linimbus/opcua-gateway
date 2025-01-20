package main

type NodeData struct {
	Node  NodeInfo
	Value *NodeValue
}

type NodeCache struct {
	Cache map[string][]NodeData
}

func NewNodeCache() *NodeCache {
	return &NodeCache{
		Cache: make(map[string][]NodeData)}
}

func (nc *NodeCache) NodeDatasGet(key string) ([]NodeData, bool) {
	node, b := nc.Cache[key]
	return node, b
}

func (nc *NodeCache) DataDatesSet(key string, data []NodeData) {
	nc.Cache[key] = data
}
