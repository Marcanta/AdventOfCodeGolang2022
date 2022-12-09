package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func convertStringArrayToIntArray(numberInString []string) []int {
	var array []int
	for _, number := range numberInString {
		integer, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		array = append(array, integer)
	}

	return array
}

func getFullyContainsPairs(limits1 []int, limits2 []int) bool {
	if !((limits1[0] <= limits2[0]) && (limits1[1] >= limits2[1])) {
		return (limits2[0] <= limits1[0]) && (limits2[1] >= limits1[1])
	}
	return true
}

// w1 - w2 < max - min no overlaps
func getOverlapsPairs(limits1 []int, limits2 []int) bool {
	width1 := limits1[1] - limits1[0]
	width2 := limits2[1] - limits2[0]
	var max int
	if limits1[1] > limits2[1] {
		max = limits1[1]
	} else {
		max = limits2[1]
	}
	var min int
	if limits1[0] < limits2[0] {
		min = limits1[0]
	} else {
		min = limits2[0]
	}
	return width1+width2 >= max-min
}

func main() {
	data := readFile("../inputs/04.in")

	assignments := strings.Split(data, "\n")

	var fullyContainsPair int
	var overlapsPairs int
	for _, assignment := range assignments {
		elves := strings.Split(assignment, ",")
		limits1 := convertStringArrayToIntArray(strings.Split(elves[0], "-"))
		limits2 := convertStringArrayToIntArray(strings.Split(elves[1], "-"))
		if getFullyContainsPairs(limits1, limits2) {
			fullyContainsPair++
		}
		if getOverlapsPairs(limits1, limits2) {
			overlapsPairs++
		}
	}

	fmt.Printf("fully contains: %d\n", fullyContainsPair)
	fmt.Printf("overlaps pairs: %d\n", overlapsPairs)
}
