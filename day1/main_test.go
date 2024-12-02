package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func parse() []string {
	// Read input file
	input, err := ioutil.ReadFile("test_input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	data := strings.Split(strings.TrimSpace(string(input)), "\n")
	
	return data
}

func TestPart1(t *testing.T) {
	testData := parse()

	expected := 11 // Replace with expected result
	got := part1(testData)

	if got != expected {
		t.Errorf("Part 1: got %d, expected %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	testData := parse()

	expected := 31 // Replace with expected result
	got := part2(testData)

	if got != expected {
		t.Errorf("Part 2: got %d, expected %d", got, expected)
	}
}
