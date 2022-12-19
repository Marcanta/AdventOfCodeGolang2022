// You can edit this code!
// Click here and start typing.R 4
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	Up    Direction = "U"
	Down            = "D"
	Right           = "R"
	Left            = "L"
)

type Knot struct {
	X, Y int
}

type Rope struct {
	head        Knot
	tail        []Knot
	tailVisited map[string]bool
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func moveTail(head Knot, tail Knot) (newTail Knot) {
	switch xDiff := head.X - tail.X; {
	case xDiff > 1:
		newTail = Knot{X: head.X - 1, Y: head.Y}
	case xDiff < -1:
		newTail = Knot{X: head.X + 1, Y: head.Y}
	}

	switch yDiff := head.Y - tail.Y; {
	case yDiff > 1:
		newTail = Knot{X: head.X, Y: head.Y - 1}
	case yDiff < -1:
		newTail = Knot{X: head.X, Y: head.Y + 1}
	}
	return
}

func (r *Rope) moveHead(direction Direction) {
	switch direction {
	case Up:
		r.head = Knot{X: r.head.X, Y: r.head.Y + 1}
	case Down:
		r.head = Knot{X: r.head.X, Y: r.head.Y - 1}
	case Right:
		r.head = Knot{X: r.head.X + 1, Y: r.head.Y}
	case Left:
		r.head = Knot{X: r.head.X - 1, Y: r.head.Y}
	}
}

func (r *Rope) moveRope(instruction string) {
	instructionSplitted := strings.Split(instruction, " ")
	var direction Direction = Direction(instructionSplitted[0])
	number, _ := strconv.Atoi(instructionSplitted[1])
	for i := 0; i < number; i++ {
		r.moveHead(direction)
		if math.Abs(float64(r.head.X)-float64(r.tail[0].X)) > 1 || math.Abs(float64(r.head.Y)-float64(r.tail[0].Y)) > 1 {
			r.tail[0] = moveTail(r.head, r.tail[0])
		}
		for i := 1; i < len(r.tail); i++ {
			if math.Abs(float64(r.tail[i-1].X)-float64(r.tail[i].X)) > 1 || math.Abs(float64(r.tail[i-1].Y)-float64(r.tail[i].Y)) > 1 {
				r.tail[i] = moveTail(r.tail[i-1], r.tail[i])
			}
		}
		r.tailVisited[fmt.Sprintf("%d,%d", r.tail[len(r.tail)-1].X, r.tail[len(r.tail)-1].Y)] = true
	}
}

func main() {
	rawData := readFile("../../inputs/09.in")

	data := strings.Split(rawData, "\n")

	// Part 1
	shortRope := Rope{head: Knot{X: 0, Y: 0}, tail: []Knot{{X: 0, Y: 0}}, tailVisited: map[string]bool{"0,0": true}}
	for _, instruction := range data {
		shortRope.moveRope(instruction)
	}

	fmt.Printf("ShortRope: , nbVisited: %d", len(shortRope.tailVisited))

	// Part 2
	longRope := Rope{head: Knot{X: 0, Y: 0}, tail: make([]Knot, 10), tailVisited: map[string]bool{"0,0": true}}
	for _, instruction := range data {
		longRope.moveRope(instruction)
	}

	fmt.Printf("rope:, nbVisited: %d", len(longRope.tailVisited))

}
