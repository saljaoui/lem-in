package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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
	if len(pathes) < len(pathe) {
		return pathe
	}
	return pathes
}

func findAllPathes(arrays [][]string, start, end string) [][]string {
	if len(arrays) == 0 {
		return nil
	}

	return find(arrays, start, end)
}

func findShoresPathes(arrays [][]string, start, end string) [][]string {
	if len(arrays) == 0 {
		return nil
	}

	sort.Slice(arrays, func(i, j int) bool {
		return len(arrays[i]) < len(arrays[j])
	})

	return find(arrays, start, end)
}

func find(arrays [][]string, start, end string) [][]string {
	result := [][]string{arrays[0]}
	seen := make(map[string]bool)

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
	graph := NewGraph()
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	file := os.Args[1]
	s, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}
	var ants int
	str := strings.Split(string(s), "\n")
	var start, end string
	for i, v := range str {
		if i == 0 {
			ants, err = strconv.Atoi(v)
			if err != nil || ants < 1 {
				fmt.Println("ERROR: invalid data format")
				return
			}
		}
		if v == "##start" {
			if i+1 < len(str) && len(str[i+1]) > 0 {
				start = string(str[i+1][0])
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
		} else if v == "##end" {
			if i+1 < len(str) && len(str[i+1]) > 0 {
				end = string(str[i+1][0])
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
		}
		st := strings.Split(v, "-")
		if len(st) == 2 {
			graph.AddEdge(st[0], st[1])
		}
	}
	if start == "" || end == "" {
		fmt.Println("ERROR: invalid data format")
		return
	}
	prev := graph.DFS(start, end)
	fmt.Println(prev)
	// printOutput(prev, ants)
}

// func printOutput(paths [][]string, ant int) {
// 	var ants = make([]string,ant) 
// 	for i := 0; i <= len(paths); i++ {
// 		for i, v := range ants {
// 			if v != "" {
// 			fmt.Print("L", i+1,"-",v)
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
