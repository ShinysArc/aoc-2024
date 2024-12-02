package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Determine input file
	inputFile := "input.txt"

	// Read input file
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	data := strings.Split(strings.TrimSpace(string(input)), "\n")

	// Solve parts
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func part1(data []string) int {
	// Implement Part 1
	return 0
}

func part2(data []string) int {
	// Implement Part 2
	return 0
}
