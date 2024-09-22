package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	funcs "lem-in/funcs"
)

func main() {
	antGraph := funcs.NewAntGraph()
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	inputFile := os.Args[1]
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}
	var antCount int
	lines := strings.Split(string(fileContent), "\n")
	
	for i, line := range lines {
		if i == 0 {
			antCount, err = strconv.Atoi(line)

			if err != nil || antCount < 1 {
				fmt.Println("ERROR: invalid data format, invalid number of Ants")
				return
			}
			continue
		}
		if line == "##start" {
			if i+1 < len(lines) && len(lines[i+1]) > 0 {
				roomInfo := strings.Split(lines[i+1], " ")
				if len(roomInfo) == 3 {
					antGraph.StartRoom = roomInfo[0]
				}
			}
		} else if line == "##end" {
			if i+1 < len(lines) && len(lines[i+1]) > 0 {
				roomInfo := strings.Split(lines[i+1], " ")
				if len(roomInfo) == 3 {
					antGraph.EndRoom = roomInfo[0]
				}
			}
		}

		connection := strings.Split(line, "-")
		if len(connection) == 2 {
			antGraph.ConnectRooms(connection[0], connection[1])
		}
	}
	if antGraph.StartRoom == "" {
		fmt.Println("ERROR: invalid data format, no start room found")
		return
	}
	if antGraph.EndRoom == "" {
		fmt.Println("ERROR: invalid data format, no end room found")
		return
	}
	validPaths := antGraph.FindAllPaths()
	if validPaths == nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	var pathsWithoutStart [][]string
	for _, path := range validPaths {
		pathsWithoutStart = append(pathsWithoutStart, path[1:])
	}
	fmt.Println(string(fileContent))
	fmt.Println()

	funcs.SimulateAntMovement(pathsWithoutStart, antCount)
}
