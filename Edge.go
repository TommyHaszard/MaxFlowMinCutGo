package main

import (
	"fmt"
	"strconv"
)

type Edge struct {
	u            *Node
	v            *Node
	maxCapacity  int
	currFlow     int
	residualEdge *Edge
	pathFlow     int
}

func NewEdge(u *Node, v *Node, maxFlow int) (*Edge, error) {
	return &Edge{u: u, v: v, maxCapacity: maxFlow, currFlow: 0, pathFlow: 0}, nil
}

func (s *Edge) AugmentFlow(flow int) {
	s.currFlow += flow
	s.residualEdge.currFlow -= flow
	s.pathFlow = flow
}

func (s *Edge) GetRemainingCapacity() int {
	return s.maxCapacity - s.currFlow
}

func (s *Edge) SetResidualEdge(residual *Edge) {
	s.residualEdge = residual
}

func (s *Edge) PrintEdgeInformation() {
	if s.maxCapacity == 0 {
		fmt.Println("From: " + s.u.name + " to: " + s.v.name + ", remaining: " + strconv.FormatInt(int64(s.GetRemainingCapacity()), 10) + "/" + strconv.FormatInt(int64(s.maxCapacity), 10) + ", curr flow " + strconv.FormatInt(int64(s.currFlow), 10) + ", minFlow: " + strconv.FormatInt(int64(s.pathFlow), 10) + ", ooo augmenting path")
	} else {
		fmt.Println("From: " + s.u.name + " to: " + s.v.name + ", remaining: " + strconv.FormatInt(int64(s.GetRemainingCapacity()), 10) + "/" + strconv.FormatInt(int64(s.maxCapacity), 10) + ", curr flow " + strconv.FormatInt(int64(s.currFlow), 10) + ", minFlow: " + strconv.FormatInt(int64(s.pathFlow), 10))
	}
}
