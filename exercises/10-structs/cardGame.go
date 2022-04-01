package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	rank string // Ace, Queen, 9, etc.
	suit rune   // Diamond, Spade, etc
}

func (c Card) String() string {
	return fmt.Sprintf("%s%c", c.rank, c.suit)
}

type Deck []Card
type Hand []Card

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

func main() {
	// Create a fill a deck
	deck := Deck{}
	deck.Init()
	fmt.Println("Original Deck: ")
	fmt.Println(deck)
	// Shuffle the deck
	deck.Shuffle()
	fmt.Println("Shuffled Deck: ")
	fmt.Println(deck)
	// Deal a hand of cards
	hand := deck.DealHand()
	fmt.Println("Hand: ")
	fmt.Println(hand)
	// Output the smaller deck
	fmt.Println("Modified Deck: ")
	fmt.Println(deck)
}
