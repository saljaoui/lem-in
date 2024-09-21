package funcs

func (g *AntGraph) FindAllPaths(start, end string) [][]string {
	var allPaths [][]string
	visited := make(map[string]bool)
	currentPath := []string{start}

	var DepthFirstSearch func(node string)
	DepthFirstSearch = func(node string) {
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
				DepthFirstSearch(neighbor)
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
		visited[node] = false
	}

	DepthFirstSearch(start)

	allUniquePaths := FindUniquePaths(allPaths, start, end)
	shortestUniquePaths := FindShortestUniquePaths(allPaths, start, end)

	if len(shortestUniquePaths) < len(allUniquePaths) {
		return allUniquePaths
	}

	return shortestUniquePaths
}
