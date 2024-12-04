package main

import (
	"fmt"

	"adventofcode2024/utils"
)

func main() {
	// Read input file
	data := utils.ParseInput("input.txt")

	// Solve parts
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func checkPartOne(array []string, i, j, dRow, dCol int) bool {
	rows := len(array)
	cols := len(array[0])

	for step := 1; step <= 3; step++ {
		newRow := i + step*dRow
		newCol := j + step*dCol
		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
			return false
		}
	}

	return array[i+dRow][j+dCol] == 'M' &&
		array[i+2*dRow][j+2*dCol] == 'A' &&
		array[i+3*dRow][j+3*dCol] == 'S'
}

func part1(data []string) int {
	directions := [][2]int{
		{-1, 0},  // Up
		{-1, -1}, // Up-left diagonal
		{0, -1},  // Left
		{1, -1},  // Down-left diagonal
		{1, 0},   // Down
		{1, 1},   // Down-right diagonal
		{0, 1},   // Right
		{-1, 1},  // Up-right diagonal
	}

	res := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] == 'X' {
				for _, dir := range directions {
					dRow, dCol := dir[0], dir[1]
					if checkPartOne(data, i, j, dRow, dCol) {
						res++
					}
				}
			}
		}
	}
	return res
}

func checkPartTwo(array []string, i, j int) bool {
	rows := len(array)
	cols := len(array[0])

	if !(i-1 >= 0 && j-1 >= 0 && i+1 < rows && j+1 < cols) {
		return false
	}
	if !(i+1 < rows && j-1 >= 0 && i-1 >= 0 && j+1 < cols) {
		return false
	}

	topLeftToBottomRight := (array[i-1][j-1] == 'M' && array[i][j] == 'A' && array[i+1][j+1] == 'S') ||
		(array[i+1][j+1] == 'M' && array[i][j] == 'A' && array[i-1][j-1] == 'S')

	topRightToBottomLeft := (array[i-1][j+1] == 'M' && array[i][j] == 'A' && array[i+1][j-1] == 'S') ||
		(array[i+1][j-1] == 'M' && array[i][j] == 'A' && array[i-1][j+1] == 'S')

	return topLeftToBottomRight && topRightToBottomLeft
}

func part2(data []string) int {
	res := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] == 'A' {
				if checkPartTwo(data, i, j) {
					res++
				}
			}
		}
	}

	return res
}
