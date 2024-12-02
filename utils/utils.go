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