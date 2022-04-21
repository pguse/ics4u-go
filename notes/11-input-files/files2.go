package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var colorRed = "\033[31m"
var colorReset = "\033[0m"

func main() {
	data, err := ioutil.ReadFile("fiveLetterWords.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\r\n")
	fmt.Println(lines)
}
