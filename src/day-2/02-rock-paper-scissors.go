package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	// winLoseMap := map[string]map[string]int{
	// 	"A": {"X": 4, "Y": 8, "Z": 3},
	// 	"B": {"X": 1, "Y": 5, "Z": 9},
	// 	"C": {"X": 7, "Y": 2, "Z": 6},
	// }
	winLoseMap := map[string]map[string]int{
		"A": {"X": 3, "Y": 4, "Z": 8},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 2, "Y": 6, "Z": 7},
	}

	data := readFile("../inputs/02.in")

	rounds := strings.Split(data, "\n")

	var sum int
	for _, round := range rounds {
		actions := strings.Split(round, " ")
		sum += winLoseMap[actions[0]][actions[1]]
	}

	fmt.Printf("sum: %d", sum)
}
