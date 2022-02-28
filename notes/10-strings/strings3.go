package main

import "fmt"

func main() {
	text := "Hello 🐱!"

	for i := 0; i < len(text); i++ {
		fmt.Printf("%d %v\n", i, text[i])
	}

	for i := 0; i < len(text); i++ {
		fmt.Printf("%d %c\n", i, text[i])
	}
}
