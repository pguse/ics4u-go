package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("fiveLetterWords.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\r\n")
	//fmt.Println(lines)
	fmt.Printf("%T", lines)
	//fmt.Println(string(data))
}
