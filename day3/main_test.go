package main

import (
	"testing"

	"adventofcode2024/utils"
)

func TestPart1(t *testing.T) {
	testData := utils.ParseInput("test_input.txt")

	expected := 161 // Replace with expected result
	got := part1(testData)

	if got != expected {
		t.Errorf("Part 1: got %d, expected %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	testData := utils.ParseInput("test_input_2.txt")

	expected := 48 // Replace with expected result
	got := part2(testData)

	if got != expected {
		t.Errorf("Part 2: got %d, expected %d", got, expected)
	}
}
