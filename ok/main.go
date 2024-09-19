package main

import (
	"fmt"
	"sort"
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

				if i < len(ants)-1 && ants[i+1].position < len(ants[i+1].path)-1 {
					ants[i+1].position++
					moves = append(moves, fmt.Sprintf("L%d-%s", ants[i+1].id, ants[i+1].path[ants[i+1].position]))
				}
				
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

for ant := 1; ant <= antCount; ant++ {
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i].rooms)+paths[i].ants < len(paths[j].rooms)+paths[j].ants
		})
		paths[0].ants++
	}

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i].rooms) < len(paths[j].rooms)
	})

p := 0
	for antIndex < antCount {
if p == 0 {
	paths[p].ants++
} else {
	paths[p].ants++
	ants[antIndex] = Ant{id: antIndex + 1, path: paths[p].rooms, position: -1}
	antIndex++
	p = 0
	continue
}

		ants[antIndex] = Ant{id: antIndex + 1, path: paths[p].rooms, position: -1}
		antIndex++
		p = 1
	}

	return ants
}

func main() {
	paths := [][]string{
		{"3"},
		{"1", "2", "3"},
	}
	antCount := 4

	simulateAntMovement(paths, antCount)
}