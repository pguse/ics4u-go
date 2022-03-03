package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("emma.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// MODIFY THE CODE BELOW
	// TO ADD THE NUMBER OF VOWELS a,e,i,o,u
	//
	sum := 0
	for _, ch := range data {
		if ch == 'a' {
			sum += 1
		}
	}
	fmt.Printf("Number of a's: %d", sum)
}
