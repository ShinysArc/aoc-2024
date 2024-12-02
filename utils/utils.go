package utils

import (
	"fmt"
	"io/ioutil"
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