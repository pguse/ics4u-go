package main

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
