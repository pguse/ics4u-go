package main

import "fmt"

func main() {
	text := "Hello 🐱!"

	for i, ch := range text {
		fmt.Printf("%d %c\n", i, ch)
	}

	fmt.Println()

	for _, ch := range text {
		fmt.Printf("%c", ch)
	}
}
