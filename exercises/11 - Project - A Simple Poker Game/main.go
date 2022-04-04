package main

import (
	"fmt"
)

func main() {
	c := Card{"A", '♣'}
	fmt.Println("Welcome to our Poker Game")
	fmt.Printf("Here is a card: %s\n", c)
}
