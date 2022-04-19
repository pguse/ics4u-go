package main

import (
	"fmt"
)

func test() {
	// A hand with nothing
	h1 := Hand{Card{"A", '♢'}, Card{"3", '♡'}, Card{"7", '♣'}, Card{"J", '♠'}, Card{"K", '♡'}}
	// A hand with a flush
	h2 := Hand{Card{"A", '♣'}, Card{"3", '♣'}, Card{"5", '♣'}, Card{"7", '♣'}, Card{"9", '♣'}}

	// Testing Flush()
	if !h1.Flush() && h2.Flush() {
		fmt.Println("Flush: OK")
	} else {
		fmt.Println("Flush: Not OK")
	}

	// Testing Straight()

	// Testing StraightFlush()

	// Testing Pair()

	// Testing ThreeOfAKind()

	// Testing FourOfAKind()

	// Testing FullHouse()

}
