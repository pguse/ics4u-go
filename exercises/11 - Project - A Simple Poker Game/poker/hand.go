package main

var value = map[string]int{
	"A":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
}

type Hand []Card

func (h Hand) String() string {
	return ""
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
