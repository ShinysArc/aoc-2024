package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"adventofcode2024/utils"
)

type order struct {
	before map[string]map[string]struct{}
}

func main() {
	// Read input file
	data := utils.ParseInput("input.txt")

	// Solve parts
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func parseOrderRules(ruleLines []string) *order {
	o := &order{
		before: make(map[string]map[string]struct{}),
	}
	for _, line := range ruleLines {
		parts := strings.Split(line, "|")
		prev, next := parts[0], parts[1]

		if _, found := o.before[next]; !found {
			o.before[next] = make(map[string]struct{})
		}
		o.before[next][prev] = struct{}{}
	}
	return o
}

func (o *order) correctOrder(update []string) bool {
	for i, u1 := range update {
		for _, u2 := range update[i+1:] {
			if _, required := o.before[u1][u2]; required {
				return false
			}
		}
	}
	return true
}

func (o *order) reorder(update []string) []string {
	slices.SortFunc(update, func(a, b string) int {
		if _, need := o.before[b][a]; need {
			return -1
		}
		if _, need := o.before[a][b]; need {
			return 1
		}
		return 0
	})
	return update
}

func parseUpdates(lines []string) [][]string {
	var updates [][]string
	for _, line := range lines {
		updates = append(updates, strings.Split(line, ","))
	}
	return updates
}

func part1(data []string) int {
	blocks := utils.SplitByEmptyLine(data)
	rules := parseOrderRules(blocks[0])

	res := 0
	for _, update := range parseUpdates(blocks[1]) {
		if rules.correctOrder(update) {
			num, _ := strconv.Atoi(update[len(update)/2])
			res += num
		}
	}
	return res
}

func part2(data []string) int {
	blocks := utils.SplitByEmptyLine(data)
	rules := parseOrderRules(blocks[0])

	res := 0
	for _, update := range parseUpdates(blocks[1]) {
		if !rules.correctOrder(update) {
			update = rules.reorder(update)
			num, _ := strconv.Atoi(update[len(update)/2])
			res += num
		}
	}
	return res
}
