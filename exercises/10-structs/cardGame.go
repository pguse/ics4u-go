package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A global variable that can be used in any method/function
// Maps the rank of a card to its value
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

// Card type
type Card struct {
	rank string // Ace, Queen, 9, etc.
	suit rune   // Diamond, Spade, etc
}

func (c Card) String() string {
	return fmt.Sprintf("%s%c", c.rank, c.suit)
}

// Deck type
type Deck []Card

func (d *Deck) Init() {
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []rune{'♣', '♢', '♡', '♠'}
	for _, s := range suits {
		for _, r := range ranks {
			*d = append(*d, Card{r, s})
		}
	}
}
func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		n := rand.Intn(len(d))
		swap := d[i]
		d[i] = d[n]
		d[n] = swap
	}
}

func (d *Deck) DealHand() Hand {
	h := Hand{}
	for i := 0; i < 5; i++ {
		h = append(h, (*d)[0])
		*d = (*d)[1:]
	}
	return h
}

// Hand type
type Hand []Card

func (h Hand) Flush() bool {
	// All cards have the same suit
	result := true
	for i := 0; i < len(h)-1; i++ {
		if h[i].suit != h[i+1].suit { // compare the suit of each card to the first card's suit
			result = false
			break
		}
	}
	return result
}

func (h Hand) Straight() bool {
	// Cards are in sequential order - assume a sorted hand
	result := true
	for i := 0; i < len(h)-1; i++ {
		if value[h[i].rank]+1 != value[h[i+1].rank] {
			result = false
			break
		}
	}
	return result
}

func (h Hand) StraightFlush() bool {
	// Cards are in sequential order with same suit - assume a sorted hand
	if h.Straight() && h.Flush() {
		return true
	} else {
		return false
	}

}

func (h Hand) Pair() bool {
	// Two cards with the same rank
	return false
}

func main() {
	// Create a fill a deck
	deck := Deck{}
	deck.Init()
	/*fmt.Println("\nOriginal Deck: ")
	fmt.Println(deck)
	// Shuffle the deck*/
	//deck.Shuffle()
	//fmt.Println("Shuffled Deck: ")
	//fmt.Println(deck)
	// Deal a hand of cards
	hand := deck.DealHand()
	fmt.Println(hand)
	//fmt.Println("Hand: ")
	fmt.Println(hand.StraightFlush())
	/*fmt.Println(hand)
	// Output the smaller deck
	fmt.Println("Modified Deck: ")
	fmt.Println(deck)*/
}
