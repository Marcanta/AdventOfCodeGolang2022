package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elves struct {
	total int
}

type ByCalories []Elves

func (a ByCalories) Len() int           { return len(a) }
func (a ByCalories) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCalories) Less(i, j int) bool { return a[i].total > a[j].total }

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func parseElves(rawElves []string) []Elves {
	var elves = []Elves{}

	for _, rawElve := range rawElves {
		var sum int
		calories := strings.Split(rawElve, "\n")
		for _, calorie := range calories {
			number, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		}
		elves = append(elves, Elves{
			total: sum,
		})
	}

	return elves
}

func getFirstNElvesAndSum(elves []Elves, N int) ([]Elves, int) {
	var sum int
	var firstElves []Elves
	for i := 0; i < N; i++ {
		fmt.Printf("elve: %d, calories: %d\n", elves[i], elves[i].total)
		sum += elves[i].total
		firstElves = append(firstElves, elves[i])
	}
	return firstElves, sum
}

func main() {
	fileContent := readFile("../inputs/01.in")

	rawElves := strings.Split(fileContent, "\n\n")

	elves := parseElves(rawElves)

	sort.Sort(ByCalories(elves))

	firstElves, sum := getFirstNElvesAndSum(elves, 3)

	fmt.Printf("sum: %d, 3 elves: %s\n", sum, firstElves)
}
