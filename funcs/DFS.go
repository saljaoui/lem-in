package funcs

import "fmt"

func (g *AntGraph) Dfs() [][]string {
	var allPaths [][]string
	visited := make(map[string]bool)
	currentPath := []string{g.StartRoom}

	var DepthFirstSearch func(node string)
	DepthFirstSearch = func(node string) {
		if node == g.EndRoom {
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			allPaths = append(allPaths, pathCopy)
			return
		}

		visited[node] = true
		for _, neighbor := range g.connections[node] {
			if !visited[neighbor] {
				currentPath = append(currentPath, neighbor)
				DepthFirstSearch(neighbor)
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
		visited[node] = false
	}

	DepthFirstSearch(g.StartRoom)
	if len(allPaths) == 0 {
		return nil
	}
	allUniquePaths := g.FilterUniquePaths(allPaths)
	shortestUniquePaths := g.FindShortestUniquePaths(allPaths)
	fmt.Println(g.Ants)
	if len(allUniquePaths) > g.Ants {
		return shortestUniquePaths
	}
	if len(shortestUniquePaths) < len(allUniquePaths) {
		return allUniquePaths
	}

	return shortestUniquePaths
}
