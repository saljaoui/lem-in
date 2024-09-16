package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Graph represents a graph using an adjacency list.
type Graph struct {
	adjacencyList map[string][]string
}

// NewGraph creates a new graph.
func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[string][]string)}
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(from, to string) {
	g.adjacencyList[from] = append(g.adjacencyList[from], to)
	g.adjacencyList[to] = append(g.adjacencyList[to], from) // For undirected graph
}

// BFS performs Breadth-First Search starting from the given start node to find the shortest path to end node.
func (g *Graph) BFS(start, end string) []string {
	visited := make(map[string]bool)
	queue := []string{}
	parent := make(map[string]string) // Maps node to its predecessor
	var paths []string

	// Enqueue the start node and mark it as visited.
	queue = append(queue, start)
	visited[start] = true
	parent[start] = "" // Start node has no parent

	for len(queue) > 0 {
		// Dequeue a node from the front of the queue.
		node := queue[0]
		queue = queue[1:]

		// If we've reached the end node, reconstruct and return the path.
		if node == end {
			path := reconstructPath(parent, end)
			slice := g.adjacencyList[start]
			newSlice := []string{}
			for _, value := range slice {
				if value != path[1] {
					newSlice = append(newSlice, value)
				}
			}
			g.adjacencyList[start] = newSlice
			// delete(g.adjacencyList, "0")

			var res string
			for _, s := range path {
				res += s
			}

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

		// Enqueue all unvisited neighbors.
		for _, neighbor := range g.adjacencyList[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				parent[neighbor] = node
			}
		}
	}

	// Return an empty path if no path is found
	return []string{}
}

// reconstructPath reconstructs the path from start to end using the parent map.
func reconstructPath(parent map[string]string, end string) []string {
	var path []string
	for node := end; node != ""; node = parent[node] {
		path = append([]string{node}, path...) // Prepend node to path
	}
	return path
}

// AddRoom adds a new room to the graph with no edges.
func (g *Graph) AddRoom(room string) {
	if _, exists := g.adjacencyList[room]; !exists {
		g.adjacencyList[room] = nil
	}
}

func main() {
	// Create a new graph and add edges.
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
		} else if s ==  "##start" {
			start = edges[i+1]
		} else if s == "##end" {
			end = edges[i+1]
		}
	}

	// Perform BFS starting from node "0" to find the path to node "5".
	fmt.Println("Shortest path from node 0 to node 5:")
	fmt.Println(start)
	fmt.Println(end)
	path := graph.BFS(start, end)
	_ = path
}
