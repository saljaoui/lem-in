package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}