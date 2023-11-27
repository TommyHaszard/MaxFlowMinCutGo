package main

type Node struct {
	name     string
	edgeList []*Edge
}

func NewNode(name string) (*Node, error) {
	return &Node{name: name, edgeList: make([]*Edge, 0)}, nil
}

func (s *Node) AddEdge(edge *Edge) {
	s.edgeList = append(s.edgeList, edge)
}

func (s *Node) GetEdges() *[]*Edge {
	return &s.edgeList
}
