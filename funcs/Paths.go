package funcs

import "sort"

func NewAntGraph() *AntGraph {
	return &AntGraph{connections: make(map[string][]string)}
}

func (g *AntGraph) ConnectRooms(room1, room2 string) {
	g.connections[room1] = append(g.connections[room1], room2)
	g.connections[room2] = append(g.connections[room2], room1) // For undirected graph
}

func FindUniquePaths(paths [][]string, start, end string) [][]string {
	if len(paths) == 0 {
		return nil
	}

	return FilterUniquePaths(paths, start, end)
}

func FindShortestUniquePaths(paths [][]string, start, end string) [][]string {
	if len(paths) == 0 {
		return nil
	}

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	return FilterUniquePaths(paths, start, end)
}

func FilterUniquePaths(paths [][]string, start, end string) [][]string {
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
