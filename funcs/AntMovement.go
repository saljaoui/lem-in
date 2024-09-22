package funcs

import (
	"fmt"
	"sort"
	"strings"
)

func SimulateAntMovement(availablePaths [][]string, antCount int) {
	paths := make([]Path, len(availablePaths))
	for i, rooms := range availablePaths {
		paths[i] = Path{id: i + 1, rooms: rooms, ants: 0}
	}

	ants := DistributeAntsToPath(paths, antCount)

	maxSteps := 0
	for _, path := range availablePaths {
		if len(path) > maxSteps {
			maxSteps = len(path)
		}
	}

	for step := 0; step < maxSteps+antCount; step++ {
		isEnd := false
		roomOccupancy := make(map[string]bool)
		pathUsed := make(map[string]bool)
		var moves []string

		for i := range ants {
			if ants[i].position < len(ants[i].path)-1 {
				isEnd = true
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
		if !isEnd {
			break
		}
	}
}

func DistributeAntsToPath(paths []Path, antCount int) []Ant {
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
