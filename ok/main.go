package main

import (
	"fmt"
)

// Function to check if a string has no repeated characters
func hasNoRepeatedChars(str string) bool {
	charSet := make(map[rune]bool)
	for _, char := range str {
		if charSet[char] {
			return false // Character is repeated
		}
		charSet[char] = true
	}
	return true // No repeated characters
}

// Filter function to keep only strings with unique characters
func filterUniqueStrings(strings []string) []string {
	var filtered []string
	for _, str := range strings {
		if hasNoRepeatedChars(str) {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func main() {
	array := []string{"abcd", "aabb", "xyz", "123", "1k"}
	filteredArray := filterUniqueStrings(array)
	fmt.Println(filteredArray) 
}
