package main

import (
	"fmt"
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

func safe(list []int) (bool, int) {
	isInc := list[0] <= list[1]

	//fmt.Println(list)
	for i := 1; i < len(list); i++ {
		var diff int
		if isInc {
			diff = list[i] - list[i-1]
		} else {
			diff = list[i-1] - list[i]
		}
		if diff < 1 || diff > 3 {
			return false, i
		}
	}

	return true, 0
}

func part1(data []string) int {
	var res int
	for _, line := range data {
		var list []int
		strList := strings.Fields(line)
		for _, el := range strList {
			list = append(list, utils.ToInt(el))
		}
		isSafe, _ := safe(list)
		if isSafe {
			res++
		}
	}

	return res
}

func remove(slice []int, s int) []int {
	newSlice := make([]int, 0, len(slice)-1)
    newSlice = append(newSlice, slice[:s]...)
    newSlice = append(newSlice, slice[s+1:]...)
    return newSlice
}

func part2(data []string) int {
	var res int
	for _, line := range data {
		var list []int
		strList := strings.Fields(line)
		for _, el := range strList {
			list = append(list, utils.ToInt(el))
		}
		isSafe, _ := safe(list)
		if isSafe {
			res++
		} else {
			for i := 0; i < len(list); i++ {
				list1 := remove(list, i)
				isSafe, _ = safe(list1)
				if isSafe {
					res++
					break
				}
			}
		}
	}

	return res
}
