package main

import "fmt"

type Graph struct {
	Vertices map[string]*Vertex
}

type Vertex struct {
	Val   string
	Edges map[string]*Edge
}

type Edge struct {
	Vertex *Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
	}
}

func (g *Graph) AddVertex(key, val string) {
	g.Vertices[key] = &Vertex{Val: val, Edges: make(map[string]*Edge)}
}

func (g *Graph) AddEdge(srcKey, destKey string) {
	// check if src & dest exist
	if _, ok := g.Vertices[srcKey]; !ok {
		return
	}
	if _, ok := g.Vertices[destKey]; !ok {
		return
	}

	// add edge src --> dest
	g.Vertices[srcKey].Edges[destKey] = &Edge{Vertex: g.Vertices[destKey]}
}

func (g *Graph) Neighbors(srcKey string) []string {
	result := []string{}
	
	if vertex, ok := g.Vertices[srcKey]; ok {
		for _, edge := range vertex.Edges {
			fmt.Println(edge.Vertex)
			result = append(result, edge.Vertex.Val)
		}
	}

	return result
}

type GraphOption func(g *Graph)

func WithAdjacencyList(list map[string][]string) GraphOption {
	return func(g *Graph) {
		for vertex, edges := range list {
			// add vertex
			if _, ok := g.Vertices[vertex]; !ok {
				g.AddVertex(vertex, vertex)
			}

			// add edges to vertex
			for _, edge := range edges {
				// add edge as vertex, if not added
				if _, ok := g.Vertices[edge]; !ok {
					g.AddVertex(edge, edge)
				}

				g.AddEdge(vertex, edge) // no weights in this adjacency list
			}
		}
	}
}

func main() {
	adjacencyList := map[string][]string{
		"1": {"2", "4"},
		"2": {"3", "5", "1"},
		"3": {"6", "2"},
		"4": {"1", "5", "7"},
		"5": {"2", "6", "8", "4"},
		"6": {"3", "0", "9", "5"},
		"7": {"4", "8"},
		"8": {"5", "9", "7"},
		"9": {"6", "0", "8"},
	}

	graph := NewGraph()
	WithAdjacencyList(adjacencyList)(graph)

	// Example usage
	fmt.Println("Neighbors of vertex '5':", graph.Neighbors("5"))
}