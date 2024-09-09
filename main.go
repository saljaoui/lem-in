package main

import (
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}