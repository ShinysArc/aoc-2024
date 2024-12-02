package main

import (
	"fmt"
	"io/ioutil"
    "math"
    "sort"
    "strconv"
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
    var list1, list2 []int

	// Parse input and populate the lists
	for _, line := range data {
		el := strings.Fields(line)
		first, _ := strconv.Atoi(el[0])
		second, _ := strconv.Atoi(el[1])
		list1 = append(list1, first)
		list2 = append(list2, second)
	}

	// Sort both lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate the result
	res := 0
	for i := range list1 {
		res += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return res
}

func part2(data []string) int {
    freq := make(map[int]int)
    for _, line := range data {
        el := strings.Fields(line)
		second, _ := strconv.Atoi(el[1])
        freq[second]++;
    }

    var res int
    for _, line := range data {
        el := strings.Fields(line)
		first, _ := strconv.Atoi(el[0])
        res += first * freq[first]
    }

	return res
}
