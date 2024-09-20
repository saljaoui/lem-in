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
			continue
		}
		if v == "##start" {
			if i+1 < len(str) && len(str[i+1]) > 0 {
				start = str[i+1]
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
		} else if v == "##end" {
			if i+1 < len(str) && len(str[i+1]) > 0 {
				end = str[i+1]
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
	var res []string
	var save [][]string
	for _, s := range prev {
		res = append(res, s[1:]...)
		save = append(save, res)
		res = nil
	}

	fmt.Println(prev)
	simulateAntMovement(save, ants)
}

type Ant struct {
	id       int
	path     []string
	position int
	isEnd    bool
}

type Path struct {
	id    int
	rooms []string
	ants  int
}


func simulateAntMovement(stringPaths [][]string, antCount int) {
	paths := make([]Path, len(stringPaths))
	for i, rooms := range stringPaths {
		paths[i] = Path{id: i + 1, rooms: rooms, ants: 0}
	}

	
	ants := assignAntsToPath(paths, antCount)

	maxSteps := 0
	for _, path := range stringPaths {
		if len(path) > maxSteps {
			maxSteps = len(path)
		}
	}

	allsteps := len(stringPaths)

	for step := 0; step < maxSteps+antCount+15; step++ {
		
		var moves []string
		for i := range ants {

			 if !ants[i].isEnd && ants[i].position < len(ants[i].path)-1 && i < allsteps {
				if i == len(ants)-1 {
					fmt.Printf("hi")
				}
				ants[i].position++
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].id, ants[i].path[ants[i].position]))
			}
			

			if ants[i].position == len(ants[i].path)-1 {
				ants[i].isEnd = true
			}
		}
		

		allsteps += len(stringPaths)
		
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
	
}


func assignAntsToPath(paths []Path, antCount int) []Ant {
	ants := make([]Ant, antCount)
	antIndex := 0

	for ant := 1; ant <= antCount; ant++ {
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i].rooms)+paths[i].ants <= len(paths[j].rooms)+paths[j].ants
		})
		paths[0].ants++
	}
	fmt.Println(paths)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i].rooms) < len(paths[j].rooms)
	})

	// First, assign ants alternating between paths
	for antIndex < antCount && paths[0].ants > 0 {
		for i := range paths {
			if paths[i].ants > 0 {
				ants[antIndex] = Ant{id: antIndex + 1, path: paths[i].rooms, position: -1}
				antIndex++
				paths[i].ants--
			}
			if antIndex >= antCount {
				break
			}
		}
	}

	// Then, assign remaining ants to paths that still have ants
	for antIndex < antCount {
		for i := range paths {
			if paths[i].ants > 0 {
				ants[antIndex] = Ant{id: antIndex + 1, path: paths[i].rooms, position: -1}
				antIndex++
				paths[i].ants--
				break
			}
		}
	}
	return ants
}
