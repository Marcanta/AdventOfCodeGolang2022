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

func uniqChar(sentence string) bool {
	for _, char := range sentence {
		if strings.Count(sentence, string(char)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	data := readFile("../../inputs/06.in")

	for i := 0; i < len(data); i++ {
		if uniqChar(data[i : i+4]) {
			fmt.Printf("index: %d, substr: %s\n", i+4, data[i:i+4])
			break
		}
	}

	for i := 0; i < len(data); i++ {
		if uniqChar(data[i : i+14]) {
			fmt.Printf("index: %d, substr: %s\n", i+14, data[i:i+14])
			break
		}
	}
}
