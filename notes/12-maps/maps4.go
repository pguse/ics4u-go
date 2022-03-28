package main

import "fmt"

func main() {
	text := "👽😊🍟😊👽😺👽"

	emoji := make(map[rune]int)

	// Count emojis
	for _, em := range text {
		emoji[em] += 1
		fmt.Printf("%c ", em)
	}

	fmt.Println()

	for k := range emoji {
		fmt.Printf("%c : %d\n", k, emoji[k])
	}
}
