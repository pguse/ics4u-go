package main

import (
	"fmt"
	"sort"
)

var value = map[string]int{"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13}

type Card struct {
	Rank string // Ace, Queen, 9, etc.
	Suit rune   // Diamond, Spade, etc
}

func (c Card) String() string {
	return fmt.Sprintf("%s%c", c.Rank, c.Suit)
}

type Hand []Card

func (h Hand) String() string {
	hand := ""
	for _, c := range h {
		hand += fmt.Sprintf("%v ", c)
	}
	return hand
}

func (h Hand) Len() int {
	return len(h)
}

func (h Hand) Less(i, j int) bool {
	return value[h[i].Rank] < value[h[j].Rank]
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	h := Hand{Card{"J", '♣'}, Card{"A", '♠'}, Card{"5", '♢'}, Card{"10", '♡'}, Card{"A", '♡'}}
	fmt.Println(h)
	sort.Sort(h)
	fmt.Println(h)
}
