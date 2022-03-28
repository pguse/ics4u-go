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
	vowels := make(map[byte]int)
	for _, ch := range data {
		if ch == 'a' || ch == 'e' {
			vowels[ch] += 1
		}
	}

	fmt.Printf("Number of a's: %d\n", vowels['a'])
	fmt.Printf("Number of e's: %d", vowels['e'])
}
