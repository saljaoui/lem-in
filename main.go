package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	g.adjacencyList[to] = append(g.adjacencyList[to], from) // For undirected graph
}

func (g *Graph) DFS(start, end string) [][]string {
	var paths [][]string
	visited := make(map[string]bool)
	path := []string{start}

	var dfs func(node string)
	dfs = func(node string) {
		if node == end {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			paths = append(paths, pathCopy)
			return
		}

		visited[node] = true
		for _, neighbor := range g.adjacencyList[node] {
			if !visited[neighbor] {
				path = append(path, neighbor)
				dfs(neighbor)
				path = path[:len(path)-1]
			}
		}
		visited[node] = false
	}

	dfs(start)

	pathe := findAllPathes(paths, start, end)
	pathes := findShoresPathes(paths, start, end)
	if  len(pathes) < len(pathe) {
		return pathe
	}
	return pathes

}

func findAllPathes(arrays [][]string, start ,end string) [][]string {
	if len(arrays) == 0 {
		return nil
	}

	result := [][]string{arrays[0]}
	seen := make(map[string]bool)

	// Mark all elements in the first array as seen
	for _, elem := range arrays[0] {
		seen[elem] = true
	}

	for _, arr := range arrays[1:] {
		unique := true
		for _, elem := range arr {
			if elem != start && elem != end && seen[elem] {
				unique = false
				break
			}
		}
		if unique {
			result = append(result, arr)
			for _, elem := range arr {
				seen[elem] = true
			}
		}
	}
	return result
}


func findShoresPathes(arrays [][]string, start ,end string) [][]string {

	if len(arrays) == 0 {
		return nil
	}

	sort.Slice(arrays, func(i, j int) bool {
		return len(arrays[i]) < len(arrays[j])
	})

	result := [][]string{arrays[0]}
	seen := make(map[string]bool)

	// Mark all elements in the first array as seen
	for _, elem := range arrays[0] {
		seen[elem] = true
	}

	for _, arr := range arrays[1:] {
		unique := true
		for _, elem := range arr {
			if elem != start && elem != end && seen[elem] {
				unique = false
				break
			}
		}
		if unique {
			result = append(result, arr)
			for _, elem := range arr {
				seen[elem] = true
			}
		}
	}
	return result
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
		} else if s == "##start" {
			start = edges[i+1]
		} else if s == "##end" {
			end = edges[i+1]
		}
	}

	// Perform BFS starting from node "0" to find the path to node "5".
	fmt.Println("Shortest path from:")
	fmt.Println(start)
	fmt.Println(end)
	path := graph.DFS(start, end)
	fmt.Println(path)
}
