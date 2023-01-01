package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items        []int
	operation    func(old int) int
	throw        func(item int) (who int)
	inspectCount int
}

func makeOperation(rawOperation string) func(int) int {
	cleanRawOperation := strings.TrimSpace(rawOperation)
	var operator rune
	var operand int
	var err error
	if strings.Count(rawOperation, "old") == 1 {
		_, err = fmt.Sscanf(cleanRawOperation, "Operation: new = old %c %d", &operator, &operand)
	} else {
		_, err = fmt.Sscanf(cleanRawOperation, "Operation: new = old %c old", &operator)
	}
	if err != nil {
		log.Fatal("Can't scan this.")
	}
	return func(old int) int {
		value := operand
		if operand == 0 {
			value = old
		}
		if operator == '*' {
			return (old * value)
		}
		return (old + value)
	}
}

func makeThrow(rawThrow []string) func(item int) (who int) {
	var toDivide int
	var monkeyIndexTrue int
	var monkeyIndexFalse int
	_, err1 := fmt.Sscanf(strings.TrimSpace(rawThrow[0]), "Test: divisible by %d", &toDivide)
	_, err2 := fmt.Sscanf(strings.TrimSpace(rawThrow[1]), "If true: throw to monkey %d", &monkeyIndexTrue)
	_, err3 := fmt.Sscanf(strings.TrimSpace(rawThrow[2]), "If false: throw to monkey %d", &monkeyIndexFalse)

	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatalf("err1: %s, err2: %s, err3: %s", err1.Error(), err2.Error(), err3.Error())
	}

	return func(item int) (who int) {
		if item%toDivide == 0 {
			return monkeyIndexTrue
		} else {
			return monkeyIndexFalse
		}
	}
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func getItems(rawItems string) []int {
	var items []int
	splittedInputs := strings.Split(rawItems, ":")
	for _, rawItem := range strings.Split(splittedInputs[1], ",") {
		item, _ := strconv.Atoi(strings.TrimSpace(rawItem))
		items = append(items, item)
	}
	return items
}

func parseMonkey(rawMonkey string, limit *int) Monkey {
	splittedInputs := strings.Split(rawMonkey, "\n")
	var toDivide int
	_, err := fmt.Sscanf(strings.TrimSpace(splittedInputs[3]), "Test: divisible by %d", &toDivide)
	if err != nil {
		log.Fatal(err.Error())
	}
	*limit *= toDivide
	return Monkey{
		items:     getItems(splittedInputs[1]),
		operation: makeOperation(splittedInputs[2]),
		throw:     makeThrow(splittedInputs[3:]),
	}
}

func (m *Monkey) doRound(monkey *[]Monkey, limit int) {
	for _, item := range m.items {
		value := m.operation(item) % limit
		m.inspectCount++
		monkeyIndex := m.throw(value)
		(*monkey)[monkeyIndex].items = append((*monkey)[monkeyIndex].items, value)
		m.items = m.items[1:]
	}
}

func (m Monkey) String() string {
	return fmt.Sprintf("inspect: %d", m.inspectCount)
}

func main() {
	data := readFile("../../inputs/11.in")

	limit := 1

	rawMonkeys := strings.Split(data, "\n\n")
	monkeys := []Monkey{}
	for _, rawMonkey := range rawMonkeys {
		monkeys = append(monkeys, parseMonkey(rawMonkey, &limit))
	}

	fmt.Printf("limit: %v\n", limit)

	for i := 0; i < 20; i++ {
		for index := range monkeys {
			monkeys[index].doRound(&monkeys, limit)
		}
		if i == 0 {
			fmt.Printf("%v\n", monkeys)
		}
	}
	// sort.Slice(monkeys, func(i, j int) bool {
	// 	return monkeys[i].inspectCount > monkeys[j].inspectCount
	// })
	fmt.Printf("%v\n", monkeys)

	for i := 0; i < 1000; i++ {
		for index := range monkeys {
			monkeys[index].doRound(&monkeys, limit)
		}
	}

	fmt.Printf("%v", monkeys)

	// fmt.Printf("monkeys: %d\n", monkeys[0].inspectCount*monkeys[1].inspectCount)
}
