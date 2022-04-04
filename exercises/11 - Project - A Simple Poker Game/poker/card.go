package main

import "fmt"

type Card struct {
	Rank string // Ace, Queen, 9, etc.
	Suit rune   // Diamond, Spade, etc
}

func (c Card) String() string {
	return fmt.Sprintf("%s%c", c.Rank, c.Suit)
}
