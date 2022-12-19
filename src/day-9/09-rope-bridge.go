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

type Vertex struct {
	X, Y int
}

type Rope struct {
	head        Vertex
	tail        Vertex
	tailVisited map[string]bool
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func (r *Rope) moveTail(direction Direction) {
	switch direction {
	case Up:
		r.tail = Vertex{X: r.head.X, Y: r.head.Y - 1}
	case Down:
		r.tail = Vertex{X: r.head.X, Y: r.head.Y + 1}
	case Right:
		r.tail = Vertex{X: r.head.X - 1, Y: r.head.Y}
	case Left:
		r.tail = Vertex{X: r.head.X + 1, Y: r.head.Y}
	}
	r.tailVisited[fmt.Sprintf("%s,%s", r.tail.X, r.tail.Y)] = true
}

func (r *Rope) moveHead(direction Direction) {
	switch direction {
	case Up:
		r.head = Vertex{X: r.head.X, Y: r.head.Y + 1}
	case Down:
		r.head = Vertex{X: r.head.X, Y: r.head.Y - 1}
	case Right:
		r.head = Vertex{X: r.head.X + 1, Y: r.head.Y}
	case Left:
		r.head = Vertex{X: r.head.X - 1, Y: r.head.Y}
	}
}

func (r *Rope) moveRope(instruction string) {
	instructionSplitted := strings.Split(instruction, " ")
	var direction Direction = Direction(instructionSplitted[0])
	number, _ := strconv.Atoi(instructionSplitted[1])
	for i := 0; i < number; i++ {
		r.moveHead(direction)
		if math.Abs(float64(r.head.X)-float64(r.tail.X)) > 1 || math.Abs(float64(r.head.Y)-float64(r.tail.Y)) > 1 {
			r.moveTail(direction)
		}
	}
}

func main() {
	rawData := readFile("../../inputs/09.in")

	data := strings.Split(rawData, "\n")

	rope := Rope{head: Vertex{X: 0, Y: 0}, tail: Vertex{X: 0, Y: 0}, tailVisited: map[string]bool{"0,0": true}}
	for _, instruction := range data {
		rope.moveRope(instruction)
	}

	fmt.Printf("rope: %v, nbVisited: %d", rope, len(rope.tailVisited))
}
