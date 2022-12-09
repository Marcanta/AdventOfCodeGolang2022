package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func computeCharValue(char rune) int {
	if unicode.IsUpper(char) {
		return (int(char) - 64) + 26
	} else {
		return int(char) - 96
	}
}

func findPriorities(data string) int {
	rucksacks := strings.Split(data, "\n")

	var sum int
	for _, rucksack := range rucksacks {
		half := len(rucksack) / 2
		compartment1 := rucksack[:half]
		compartment2 := rucksack[half:]
		for _, char := range compartment1 {
			if strings.ContainsRune(compartment2, char) {
				sum += computeCharValue(char)
				break
			}
		}
	}

	return sum
}

func findBadges(data string) int {
	rucksacks := strings.Split(data, "\n")

	var sum int
	for i := 0; i < len(rucksacks); i += 3 {
		for _, char := range rucksacks[i] {
			bool1 := strings.ContainsRune(rucksacks[i+1], char)
			bool2 := strings.ContainsRune(rucksacks[i+2], char)
			if bool1 && bool2 {
				sum += computeCharValue(char)
				break
			}
		}
	}

	return sum
}

func main() {
	data := readFile("../inputs/03.in")

	fmt.Printf("Priorities: %d\n", findPriorities(data))

	fmt.Printf("Priorities of Badges: %d\n", findBadges(data))
}
