package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"adventofcode2024/utils"
)

func main() {
	// Read input file
	data := utils.ParseInput("input.txt")

	// Solve parts
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func part1(data []string) int {
	muls := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	res := 0
	for _, line := range data {
		matches := muls.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			res += x * y
		}
	}
	return res
}

func part2(data []string) int {
	muls := regexp.MustCompile(`^mul\((\d{1,3}),(\d{1,3})\)`)
	enabled := true
	res := 0

	for _, line := range data {
		for i := range len(line) {
			if strings.HasPrefix(line[i:], "do()") {
				enabled = true
			}
			if strings.HasPrefix(line[i:], "don't()") {
				enabled = false
			}

			match := muls.FindStringSubmatch(line[i:])
			if match != nil && enabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				res += x * y
			}
		}
	}

	return res
}
