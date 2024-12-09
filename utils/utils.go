package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ParseInput(inputFile string) []string {
	// Read input file
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	data := strings.Split(strings.TrimSpace(string(input)), "\n")
	
	return data
}

func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

func SplitByEmptyLine(data []string) [][]string {
	var res [][]string

	var textBlock []string
	for _, line := range data {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			textBlock = append(textBlock, line)
		} else {
			if len(textBlock) > 0 {
				res = append(res, textBlock)
			}
			textBlock = []string{}
		}
	}
	if len(textBlock) > 0 {
		res = append(res, textBlock)
	}
	return res
}