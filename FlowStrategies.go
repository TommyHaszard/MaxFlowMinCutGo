package main

import (
	"fmt"
	"math"
	"strconv"
)

type FlowStrategies interface {
	searchAlgo(graph *Graph) *Graph
	printMinCuts(graph *Graph)
}

type EdmondsKarp struct {
}

func (s *EdmondsKarp) searchAlgo(graph *Graph) *Graph {
	fmt.Println("Edmonds-Karp")
	flow := 0

	sink := graph.sink
	source := graph.source

	for {
		queue := make([]*Node, 0)
		backwardsEdges := make(map[string]*Edge)
		queue = append(queue, source)
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, currEdge := range curr.edgeList {
				v := currEdge.v
				unused := currEdge.GetRemainingCapacity()
				if unused > 0 && backwardsEdges[v.name] == nil {
					backwardsEdges[v.name] = currEdge
					if v.name == sink.name {
						break
					}
					queue = append(queue, v)
				}
			}
		}

		if backwardsEdges[sink.name] == nil {
			break
		}
		minFlow := math.MaxInt32
		currEdge := backwardsEdges[sink.name]
		sourceFound := false
		for !sourceFound {
			if currEdge.u.name == source.name {
				sourceFound = true
			}
			minFlow = min(minFlow, currEdge.GetRemainingCapacity())
			currEdge = backwardsEdges[currEdge.u.name]
		}
		if minFlow == 0 {
			break
		}
		currEdge = backwardsEdges[sink.name]
		sourceFound = false
		for !sourceFound {
			if currEdge.u.name == source.name {
				sourceFound = true
			}
			currEdge.AugmentFlow(minFlow)
			currEdge.PrintEdgeInformation()
			currEdge = backwardsEdges[currEdge.u.name]
		}
		flow += minFlow
		fmt.Printf("Curr Flow: %d\n", flow)
	}
	graph.maxFlow = flow
	return graph
}

func (s *EdmondsKarp) PrintMinCuts(graph *Graph) {
	sourceNodes := make(map[*Node]bool)
	source := graph.source
	queue := make([]*Node, 0)

	queue = append(queue, source)
	sourceNodes[source] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, currEdge := range curr.edgeList {
			v := currEdge.v
			if currEdge.GetRemainingCapacity() > 0 && !sourceNodes[v] {
				sourceNodes[v] = true
				queue = append(queue, v)
			}
		}
	}

	fmt.Println("\nMincut Results:\n")
	sinkNodes := make(map[*Node]bool)
	for _, currNode := range graph.vertices {
		if !sourceNodes[currNode] {
			sinkNodes[currNode] = true
		}
	}

	maxFlow := 0

	for currEdge, _ := range graph.edgeSet {
		if sourceNodes[currEdge.u] && sinkNodes[currEdge.v] {
			fmt.Println(currEdge.u.name + " - " + currEdge.v.name + " with capacity: " + strconv.FormatInt(int64(currEdge.maxCapacity), 10))
			maxFlow += int(currEdge.maxCapacity)
		}
	}
	fmt.Println("\nSum of Cut Edge Capacities = " + strconv.Itoa(maxFlow))
	fmt.Println("Max Flow of the Graph = " + strconv.Itoa(int(graph.maxFlow)))
}

type FordFulkerson struct {
}

func (s *FordFulkerson) searchAlgo(graph *Graph) *Graph {
	prevEdge := make(map[string]*Edge, 0)
	visited := make(map[*Node]bool)
	source := graph.source
	sink := graph.sink
	maxFlow := 0

	for dfs(graph, source, sink, prevEdge, visited) {
		flow := math.MaxInt
		for _, e := range prevEdge {
			flow = min(flow, e.GetRemainingCapacity())
		}

		for _, e := range prevEdge {
			e.AugmentFlow(flow)
			e.PrintEdgeInformation()
		}

		clear(prevEdge)
		clear(visited)
		maxFlow += flow
		fmt.Printf("Curr Flow: %d\n", maxFlow)
	}
	graph.maxFlow = maxFlow
	return graph
}

func dfs(graph *Graph, current *Node, sink *Node, prevEdge map[string]*Edge, visited map[*Node]bool) bool {
	if current.name == sink.name {
		return true
	}
	visited[current] = true
	for _, currEdge := range current.edgeList {
		neighbour := currEdge.v
		if !visited[neighbour] && currEdge.GetRemainingCapacity() > 0 {
			prevEdge[currEdge.v.name] = currEdge
			if dfs(graph, neighbour, sink, prevEdge, visited) {
				return true
			}
		}
	}
	return false
}
