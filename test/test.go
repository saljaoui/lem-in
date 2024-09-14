package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph struct {
	adjacencyList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[string][]string)}
}


func (g *Graph) AddEdge(from, to string) {
	g.adjacencyList[from] = append(g.adjacencyList[from], to)
	g.adjacencyList[to] = append(g.adjacencyList[to], from)
}


func (g *Graph) BFS(start, end string) []string {
	visited := make(map[string]bool)
	queue := []string{}
	parent := make(map[string]string) 
	var paths []string



	queue = append(queue, start)
	visited[start] = true
	parent[start] = "" 

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]


		if node == end {
		 path := reconstructPath(parent, end)
		 
		 visited = make(map[string]bool)
		 parent = make(map[string]string)
		 queue = nil
		 queue = append(queue, start)
		 visited[start] = true
		 parent[start] = ""

		paths = append(paths, path...)
		 for _, p := range paths {
			if p != end {
				visited[p] = true
			}
		 }
		 node = queue[0]
		 queue = queue[1:]
		 fmt.Println(path)
		}


		for _, neighbor := range g.adjacencyList[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				parent[neighbor] = node
			}
		}
	}


	return []string{}
}


func reconstructPath(parent map[string]string, end string) []string {
	var path []string
	for node := end; node != ""; node = parent[node] {
		path = append([]string{node}, path...) 
	}
	return path
}

func (g *Graph) AddRoom(room string) {
	if _, exists := g.adjacencyList[room]; !exists {
		g.adjacencyList[room] = nil
	}
}

func main() {

	graph := NewGraph()
	// edges := []string{"0-1", "0-3", "0-2", "2-5", "3-6", "1-4", "4-6"}

	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	var start string
	var end string
	contentStr := string(content)
	edges := strings.Split(contentStr, "\n")
	for i, s := range edges {
		parts := strings.Split(s, "-")
		if len(parts) == 2 {
			fmt.Println(parts)
			graph.AddEdge(parts[0], parts[1])
		} else if strings.HasPrefix(s, "##start") {
			start = edges[i+1]
		} else if strings.HasPrefix(s, "##end") {
			end = edges[i+1]
		}
	}

	// Perform BFS starting from node "0" to find the path to node "5".
	fmt.Println("Shortest path from node 0 to node 5:")
	fmt.Println(start[:1])
	fmt.Println(end[:1])
	path := graph.BFS(start[:1], end[:1])
_= path
}
