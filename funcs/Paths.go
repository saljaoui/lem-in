package funcs

import "sort"

func NewAntGraph() *AntGraph {
	return &AntGraph{
		connections: make(map[string][]string),
	}
}

func (g *AntGraph) ConnectRooms(room1, room2 string) {
	g.connections[room1] = append(g.connections[room1], room2)
	g.connections[room2] = append(g.connections[room2], room1) // For undirected graph
}

func (g *AntGraph) FindShortestUniquePaths(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	return g.FilterUniquePaths(paths)
}

func (g *AntGraph) FilterUniquePaths(paths [][]string) [][]string {
	uniquePaths := [][]string{paths[0]}
	seenRooms := make(map[string]bool)

	for _, room := range paths[0] {
		seenRooms[room] = true
	}

	for _, path := range paths[1:] {
		isUnique := true
		for _, room := range path {
			if room != g.StartRoom && room != g.EndRoom && seenRooms[room] {
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
