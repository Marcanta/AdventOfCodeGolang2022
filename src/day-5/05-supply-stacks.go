package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CratesMap map[int][]string

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func parseInstruction(instruction string) (nbCrate int, from int, to int) {
	splittedInstruction := strings.Split(instruction, " ")
	nbCrate, err := strconv.Atoi(splittedInstruction[1])
	if err != nil {
		log.Fatal(err)
	}
	from, err = strconv.Atoi(splittedInstruction[3])
	if err != nil {
		log.Fatal(err)
	}
	to, err = strconv.Atoi(splittedInstruction[5])
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (m CratesMap) execute(instruction string) {
	nbCrate, from, to := parseInstruction(instruction)
	for i := 0; i < nbCrate; i++ {
		m[to] = append([]string{m[from][0]}, m[to]...)
		m[from] = m[from][1:]
	}
}

func (m CratesMap) executeCrane9001(instruction string) {
	nbCrate, from, to := parseInstruction(instruction)
	crates := make([]string, len(m[from][:nbCrate]))
	copy(crates, m[from][:nbCrate])
	m[to] = append(crates, m[to]...)
	m[from] = m[from][nbCrate:]
}

func createCratesMap(rawData []string) CratesMap {
	positionMap := map[int]int{
		1:  1,
		5:  2,
		9:  3,
		13: 4,
		17: 5,
		21: 6,
		25: 7,
		29: 8,
		33: 9,
	}
	cratesMap := make(CratesMap)
	regexDigit := regexp.MustCompile(`[0-9]+`)
	for _, line := range rawData {
		if match := regexDigit.MatchString(line); match {
			break
		}
		for key, value := range positionMap {
			if rune(line[key]) != ' ' {
				letter := string(line[key])
				cratesMap[value] = append(cratesMap[value], letter)
			}
		}
	}
	return cratesMap
}

func checkWNumber(m CratesMap) bool {
	count := 0
	for _, value := range m {
		for _, letter := range value {
			if letter == "W" {
				count++
			}
		}
	}
	return count <= 5
}

func main() {
	data := readFile("../../inputs/05.in")

	dataSplitted := strings.Split(data, "\n\n")

	cratesMap := createCratesMap(strings.Split(dataSplitted[0], "\n"))
	fmt.Printf("cratesMap: %v\n", cratesMap)
	rawInstructions := dataSplitted[1]
	instructions := strings.Split(rawInstructions, "\n")

	for _, action := range instructions {
		// cratesMap.execute(action)
		// if !checkWNumber(cratesMap) {
		// 	log.Fatal("UN DE TROP")
		// 	os.Exit(1)
		// }
		cratesMap.executeCrane9001(action)
	}

	fmt.Printf("cratesMap: %v\n", cratesMap)
}
