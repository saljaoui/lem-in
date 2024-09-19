package main

import (
	"fmt"
	"strings"
)

type Ant struct {
	id     int
	path   []string
	position int
}

func simulateAntMovement(paths [][]string, antCount int) {
	ants := make([]Ant, antCount)
	for i := 0; i < antCount; i++ {
		ants[i] = Ant{id: i + 1, path: paths[i % len(paths)], position: -1}
	}

	maxSteps := 0
	for _, path := range paths {
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

func main() {
	paths := [][]string{
		{"3"},
		{"2", "3", "1"},
	}
	antCount := 5
	simulateAntMovement(paths, antCount)
}