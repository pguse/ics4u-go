package main

import (
	"math/rand"
	"time"
)

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
