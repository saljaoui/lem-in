package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X, Y int
}

type Colony struct {
	Ants     int
	Rooms    map[string]*Room
	Links    map[string][]string
	Start    string
	End      string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./lem-in <input_file>")
		return
	}

	colony, err := parseInput(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	printColony(colony)
	path := findPath(colony)
	if path == nil {
		fmt.Println("ERROR: No path found")
		return
	}

	movements := moveAnts(colony, path)
	printMovements(movements)
}

func parseInput(filename string) (*Colony, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	colony := &Colony{
		Rooms: make(map[string]*Room),
		Links: make(map[string][]string),
	}

	scanner := bufio.NewScanner(file)
	state := "ants"
	var nextRoomType string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line == "##start" {
			nextRoomType = "start"
			continue
		} else if line == "##end" {
			nextRoomType = "end"
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		switch state {
		case "ants":
			colony.Ants, err = strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("invalid number of ants")
			}
			state = "rooms"
		case "rooms":
			if strings.Contains(line, "-") {
				state = "links"
				// Process this line as a link
				parts := strings.Split(line, "-")
				if len(parts) != 2 {
					return nil, fmt.Errorf("invalid link format")
				}
				colony.Links[parts[0]] = append(colony.Links[parts[0]], parts[1])
				colony.Links[parts[1]] = append(colony.Links[parts[1]], parts[0])
			} else {
				// Process room
				parts := strings.Fields(line)
				if len(parts) != 3 {
					return nil, fmt.Errorf("invalid room format")
				}
				x, _ := strconv.Atoi(parts[1])
				y, _ := strconv.Atoi(parts[2])
				colony.Rooms[parts[0]] = &Room{Name: parts[0], X: x, Y: y}
				if nextRoomType == "start" {
					colony.Start = parts[0]
					nextRoomType = ""
				} else if nextRoomType == "end" {
					colony.End = parts[0]
					nextRoomType = ""
				}
			}
		case "links":
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid link format")
			}
			colony.Links[parts[0]] = append(colony.Links[parts[0]], parts[1])
			colony.Links[parts[1]] = append(colony.Links[parts[1]], parts[0])
		}
	}

	if colony.Start == "" || colony.End == "" {
		return nil, fmt.Errorf("start or end room not specified")
	}

	return colony, nil
}

func printColony(c *Colony) {
	fmt.Printf("%d\n", c.Ants)
	for _, room := range c.Rooms {
		fmt.Printf("%s %d %d\n", room.Name, room.X, room.Y)
	}
	for room, links := range c.Links {
		for _, link := range links {
			fmt.Printf("%s-%s\n", room, link)
		}
	}
	fmt.Println()
}

func findPath(c *Colony) []string {
	// Implement BFS here
	// This is a simplified version and doesn't handle all cases
	visited := make(map[string]bool)
	queue := [][]string{{c.Start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		room := path[len(path)-1]

		if room == c.End {
			return path
		}

		if !visited[room] {
			visited[room] = true
			for _, neighbor := range c.Links[room] {
				if !visited[neighbor] {
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, neighbor)
					queue = append(queue, newPath)
				}
			}
		}
	}

	return nil
}

func moveAnts(c *Colony, path []string) [][]string {
	movements := [][]string{}
	antPositions := make([]int, c.Ants)

	for !allAntsAtEnd(antPositions, len(path)-1) {
		move := []string{}
		for ant := 0; ant < c.Ants; ant++ {
			if antPositions[ant] < len(path)-1 {
				antPositions[ant]++
				move = append(move, fmt.Sprintf("L%d-%s", ant+1, path[antPositions[ant]]))
			}
		}
		if len(move) > 0 {
			movements = append(movements, move)
		}
	}

	return movements
}

func allAntsAtEnd(positions []int, endIndex int) bool {
	for _, pos := range positions {
		if pos != endIndex {
			return false
		}
	}
	return true
}

func printMovements(movements [][]string) {
	for _, move := range movements {
		fmt.Println(strings.Join(move, " "))
	}
}