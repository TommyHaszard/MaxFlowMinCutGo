package main

import "fmt"

type Graph struct {
	source   *Node
	sink     *Node
	maxFlow  int
	count    int
	edgeSet  map[*Edge]bool
	vertices map[string]*Node
}

func NewGraph(source *Node, sink *Node) (*Graph, error) {
	edgeSet := make(map[*Edge]bool)
	vertices := make(map[string]*Node)
	return &Graph{source: source, sink: sink, maxFlow: 0, count: 0, edgeSet: edgeSet, vertices: vertices}, nil
}

func (s *Graph) CreateVertex(u *Node) error {
	if s.vertices[u.name] == nil {
		s.vertices[u.name] = u
		return nil
	}
	return fmt.Errorf("failed to create Node: {%s}", u.name)
}

func (s *Graph) AddEdge(u *Node, v *Node, maxFlow int) error {
	if s.vertices[u.name] != nil && s.vertices[v.name] != nil {
		newEdge, err := NewEdge(u, v, maxFlow)
		if err != nil {
			return err
		}
		residualEdge, err := NewEdge(v, u, 0)
		if err != nil {
			return err
		}

		newEdge.SetResidualEdge(residualEdge)
		residualEdge.SetResidualEdge(newEdge)
		u.AddEdge(newEdge)
		v.AddEdge(residualEdge)
		s.edgeSet[newEdge] = true
		s.edgeSet[residualEdge] = true
	}
	return nil
}

func (s *Graph) IncrementCount() {
	s.count++
}
