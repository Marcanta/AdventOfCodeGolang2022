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

func reverseInts(array []int) []int {
	res := []int{}
	for i := len(array) - 1; i >= 0; i-- {
		res = append(res, array[i])
	}
	return res
}

func getVerticalSlice(trees [][]int, y int) []int {
	var acc []int
	for i := 0; i < len(trees); i++ {
		acc = append(acc, trees[i][y])
	}
	return acc
}

func isGreaterThanOthers(number int, other []int) (greater bool, distance int) {
	for _, v := range other {
		distance++
		if number <= v {
			return false, distance
		}
	}
	return true, distance
}

func parseTrees(rawData string) (trees [][]int) {
	parseStringsToInts := func(strings []string) (ints []int) {
		for _, s := range strings {
			nb, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, nb)
		}
		return
	}

	treeLines := strings.Split(rawData, "\n")
	for _, line := range treeLines {
		trees = append(trees, parseStringsToInts(strings.Split(line, "")))
	}
	return
}

func main() {
	data := readFile("../../inputs/08.in")

	trees := parseTrees(data)

	bestScenicScore := 0
	invisibleTrees := 0
	for i := 1; i < len(trees)-1; i++ {
		for y := 1; y < len(trees[i])-1; y++ {
			visibleFromLeft, leftViewingDistance := isGreaterThanOthers(trees[i][y], reverseInts(trees[i][:y]))
			visibleFromRight, rightViewingDistance := isGreaterThanOthers(trees[i][y], trees[i][y+1:])
			vSlice := getVerticalSlice(trees, y)
			visibleFromTop, topViewingDistance := isGreaterThanOthers(trees[i][y], reverseInts(vSlice[:i]))
			visibleFromBottom, bottomViewingDistance := isGreaterThanOthers(trees[i][y], vSlice[i+1:])

			scenicScore := leftViewingDistance * rightViewingDistance * topViewingDistance * bottomViewingDistance

			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
			if !visibleFromLeft && !visibleFromRight && !visibleFromTop && !visibleFromBottom {
				invisibleTrees++
				continue
			}
		}
	}

	treesCount := len(trees) * len(trees[0])
	visibleTrees := treesCount - invisibleTrees

	fmt.Printf("visibleTrees: %v, invisible: %v, treesCount: %v\n", visibleTrees, invisibleTrees, treesCount)
	fmt.Printf("bestScenicScore: %v\n", bestScenicScore)
}
