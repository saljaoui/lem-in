package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type AntGraph struct {
	connections map[string][]string
}

func NewAntGraph() *AntGraph {
	return &AntGraph{connections: make(map[string][]string)}
}

func (g *AntGraph) ConnectRooms(room1, room2 string) {
	g.connections[room1] = append(g.connections[room1], room2)
	g.connections[room2] = append(g.connections[room2], room1) // For undirected graph
}

func (g *AntGraph) FindAllPaths(start, end string) [][]string {
	var allPaths [][]string
	visited := make(map[string]bool)
	currentPath := []string{start}

	var depthFirstSearch func(node string)
	depthFirstSearch = func(node string) {
		if node == end {
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			allPaths = append(allPaths, pathCopy)
			return
		}

		visited[node] = true
		for _, neighbor := range g.connections[node] {
			if !visited[neighbor] {
				currentPath = append(currentPath, neighbor)
				depthFirstSearch(neighbor)
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
		visited[node] = false
	}

	depthFirstSearch(start)

	allUniquePaths := findUniquePaths(allPaths, start, end)
	shortestUniquePaths := findShortestUniquePaths(allPaths, start, end)

	if len(shortestUniquePaths) < len(allUniquePaths) {
		return allUniquePaths
	}

	return shortestUniquePaths
}

func findUniquePaths(paths [][]string, start, end string) [][]string {
	if len(paths) == 0 {
		return nil
	}

	return filterUniquePaths(paths, start, end)
}

func findShortestUniquePaths(paths [][]string, start, end string) [][]string {
	if len(paths) == 0 {
		return nil
	}

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	return filterUniquePaths(paths, start, end)
}

func filterUniquePaths(paths [][]string, start, end string) [][]string {
	uniquePaths := [][]string{paths[0]}
	seenRooms := make(map[string]bool)

	for _, room := range paths[0] {
		seenRooms[room] = true
	}

	for _, path := range paths[1:] {
		isUnique := true
		for _, room := range path {
			if room != start && room != end && seenRooms[room] {
				isUnique = false
				break
			}
		}
		if isUnique {
			uniquePaths = append(uniquePaths, path)
			for _, room := range path {
				seenRooms[room] = true
			}
		}
	}
	return uniquePaths
}

func main() {
	antGraph := NewAntGraph()
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	inputFile := os.Args[1]
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}
	var antCount int
	lines := strings.Split(string(fileContent), "\n")
	var startRoom, endRoom string
	for i, line := range lines {
		if i == 0 {
			antCount, err = strconv.Atoi(line)

			if err != nil || antCount < 1 {
				fmt.Println("ERROR: invalid data format, invalid number of Ants")
				return
			}
			continue
		}
		if line == "##start" {
			if i+1 < len(lines) && len(lines[i+1]) > 0 {
				roomInfo := strings.Split(lines[i+1], " ")
				if len(roomInfo) == 3 {
					startRoom = roomInfo[0]
				}
			}
		} else if line == "##end" {
			if i+1 < len(lines) && len(lines[i+1]) > 0 {
				roomInfo := strings.Split(lines[i+1], " ")
				if len(roomInfo) == 3 {
					endRoom = roomInfo[0]
				}
			}
		}

		connection := strings.Split(line, "-")
		if len(connection) == 2 {
			antGraph.ConnectRooms(connection[0], connection[1])
		}
	}
	if startRoom == "" {
		fmt.Println("ERROR: invalid data format, no start room found")
		return
	}
	if endRoom == "" {
		fmt.Println("ERROR: invalid data format, no end room found")
		return
	}
	validPaths := antGraph.FindAllPaths(startRoom, endRoom)
	if validPaths == nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	var pathsWithoutStart [][]string
	for _, path := range validPaths {
		pathsWithoutStart = append(pathsWithoutStart, path[1:])
	}
	fmt.Println(string(fileContent))
	fmt.Println()

	simulateAntMovement(pathsWithoutStart, antCount)
}

type Ant struct {
	id       int
	path     []string
	position int
}

type Path struct {
	id    int
	rooms []string
	ants  int
}

func simulateAntMovement(availablePaths [][]string, antCount int) {
	paths := make([]Path, len(availablePaths))
	for i, rooms := range availablePaths {
		paths[i] = Path{id: i + 1, rooms: rooms, ants: 0}
	}

	ants := distributeAntsToPath(paths, antCount)

	maxSteps := 0
	for _, path := range availablePaths {
		if len(path) > maxSteps {
			maxSteps = len(path)
		}
	}

	for step := 0; step < maxSteps+antCount; step++ {
		roomOccupancy := make(map[string]bool)
		pathUsed := make(map[string]bool)
		var moves []string

		for i := range ants {
			if ants[i].position < len(ants[i].path)-1 {
				ants[i].position++
				if roomOccupancy[ants[i].path[ants[i].position]] && pathUsed[ants[i].path[0]] {
					ants[i].position--
					continue
				}
				pathUsed[ants[i].path[0]] = true
				roomOccupancy[ants[i].path[ants[i].position]] = true
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].id, ants[i].path[ants[i].position]))
			}
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

func distributeAntsToPath(paths []Path, antCount int) []Ant {
	ants := make([]Ant, antCount)
	antIndex := 0

	for ant := 1; ant <= antCount; ant++ {
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i].rooms) < len(paths[j].rooms)
		})
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i].rooms)+paths[i].ants < len(paths[j].rooms)+paths[j].ants
		})
		paths[0].ants++
	}

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