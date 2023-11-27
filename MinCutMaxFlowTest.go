package main

import (
	"log"
)

func main() {
	graph := CycleOnSource()
	ed := &EdmondsKarp{}
	ed.searchAlgo(graph)
	ed.PrintMinCuts(graph)

	graph = CycleOnSource()
	ff := &FordFulkerson{}
	ff.searchAlgo(graph)
}

func SingleNodeGraph() *Graph {
	source, err := NewNode("Source")
	if err != nil {
		log.Fatal(err)
	}
	sink, err := NewNode("Sink")
	if err != nil {
		log.Fatal(err)
	}

	graph, err := NewGraph(source, sink)
	if err != nil {
		log.Fatal(err)
	}

	node1, err := NewNode("Node1")
	if err != nil {
		log.Fatal(err)
	}

	graph.CreateVertex(source)
	graph.CreateVertex(node1)
	graph.CreateVertex(sink)
	graph.AddEdge(source, node1, 5)
	graph.AddEdge(node1, sink, 5)
	return graph
}

func CycleOnSource() *Graph {

	source, err := NewNode("Source")
	if err != nil {
		log.Fatal(err)
	}
	sink, err := NewNode("Sink")
	if err != nil {
		log.Fatal(err)
	}
	graph, err := NewGraph(source, sink)
	if err != nil {
		log.Fatal(err)
	}
	node1, err := NewNode("Node1")
	if err != nil {
		log.Fatal(err)
	}
	node2, err := NewNode("Node2")
	if err != nil {
		log.Fatal(err)
	}
	node3, err := NewNode("Node3")
	if err != nil {
		log.Fatal(err)
	}
	graph.CreateVertex(source)
	graph.CreateVertex(node1)
	graph.CreateVertex(node2)
	graph.CreateVertex(node3)
	graph.CreateVertex(sink)
	graph.AddEdge(source, node1, 9)
	graph.AddEdge(source, node3, 4)
	graph.AddEdge(node1, node2, 7)
	graph.AddEdge(node2, source, 2)
	graph.AddEdge(node2, node3, 2)
	graph.AddEdge(node2, sink, 5)
	graph.AddEdge(node3, sink, 6)
	return graph
}
