package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph struct {
	
}


func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	file := strings.Split(string(content), "\n")
	fmt.Println(file)
}