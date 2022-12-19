package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CRTImage string

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func computeTargetCycle(targetCycle *[]int, X int, cycle int) (signalStrength int) {
	if cycle != (*targetCycle)[0] {
		return
	}
	signalStrength += X * (*targetCycle)[0]
	copy(*targetCycle, (*targetCycle)[1:])
	// if len(*targetCycle) == 0 {
	// 	panic("end this.")
	// }
	return
}

func (i *CRTImage) drawSprite(cycle int, X int) {
	if diff := X - cycle%40; diff < 1 && diff > -3 {
		*i += "#"
	} else {
		*i += "."
	}
	if cycle%40 == 0 {
		*i += "\n"
	}
}

func main() {
	rawData := readFile("../../inputs/10.in")
	instructions := strings.Split(rawData, "\n")

	// Part 1
	targetCycle := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	cycle := 1
	X := 1
	// Part 2
	var image CRTImage = ""

	for _, order := range instructions {
		if len(targetCycle) != 0 {
			sum += computeTargetCycle(&targetCycle, X, cycle)
		}
		image.drawSprite(cycle, X)
		if order == "noop" {
			cycle++
			continue
		}
		if orderSplitted := strings.Split(order, " "); orderSplitted[0] == "addx" {
			if len(targetCycle) != 0 {
				sum += computeTargetCycle(&targetCycle, X, cycle+1)
			}
			image.drawSprite(cycle+1, X)
			cycle += 2
			toAdd, _ := strconv.Atoi(orderSplitted[1])
			X += toAdd
		}
	}

	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("%s\n", image)
}
