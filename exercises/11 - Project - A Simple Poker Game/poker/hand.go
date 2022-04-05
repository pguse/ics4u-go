package main

import "fmt"

var value = map[string]int{"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13}

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

func (h Hand) Flush() bool {
	// All cards have the same suit
	return false
}

func (h Hand) Straight() bool {
	// Cards are in sequential order - assume a sorted hand
	return false
}

func (h Hand) StraightFlush() bool {
	// Cards are in sequential order with same suit - assume a sorted hand
	return false
}

func (h Hand) Pair() bool {
	// Two cards with the same rank
	return false
}
