package main

import (
	"fmt"
	"sort"
	"strings"
)

type Path struct {
	id    int
	rooms []string
	ants  int
}

type Ant struct {
	id       int
	pathID   int
	position int
}

func assignAntsToPath(paths []Path, antCount int) []Path {
	for ant := 1; ant <= antCount; ant++ {
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i].rooms)+paths[i].ants < len(paths[j].rooms)+paths[j].ants
		})
		paths[0].ants++
	}
	return paths
}

func simulateAntMovement(stringPaths [][]string, antCount int) {
	paths := make([]Path, len(stringPaths))
	for i, rooms := range stringPaths {
		paths[i] = Path{id: i + 1, rooms: rooms, ants: 0}
	}

	assignedPaths := assignAntsToPath(paths, antCount)

	ants := make([]Ant, antCount)
	antIndex := 0
	for _, path := range assignedPaths {
		for i := 0; i < path.ants; i++ {
			ants[antIndex] = Ant{id: antIndex + 1, pathID: path.id, position: -1} // Start before first room
			antIndex++
		}
	}

	maxSteps := 0
	for _, path := range paths {
		if len(path.rooms) > maxSteps {
			maxSteps = len(path.rooms)
		}
	}

	for step := 0; step < maxSteps; step++ {
		moves := []string{}
		for i := range ants {
			if ants[i].position < len(paths[ants[i].pathID-1].rooms) - 1 {
				ants[i].position++
				currentRoom := paths[ants[i].pathID-1].rooms[ants[i].position]
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].id, currentRoom))
			}
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

func main() {
	stringPaths := [][]string{
		{"0", "3"},
		{"0", "1", "2", "3"},
	}

	antCount := 6

	simulateAntMovement(stringPaths, antCount)
}