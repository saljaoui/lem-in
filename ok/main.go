package main

import (
	"fmt"
	"strings"
)

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

	for step := 0; step < maxSteps+antCount; step++ {
		var moves []string
		for i := range ants {
			if step >= i && ants[i].position < len(ants[i].path)-1 {
				ants[i].position++
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].id, ants[i].path[ants[i].position]))
			}
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

func assignAntsToPath(paths []Path, antCount int) []Ant {
	ants := make([]Ant, antCount)
	antIndex := 0

	// First, assign ants alternating between paths
	for antIndex < antCount && paths[0].ants > 0 && paths[1].ants > 0 {
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

func main() {
	paths := [][]string{
		{"3"},
		{"1", "2", "3"},
	}
	antCount := 4 // Adjust this number as needed
	simulateAntMovement(paths, antCount)
}
