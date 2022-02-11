package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := [3][3]int{}

	// Create a random 2D array with values 0->9
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			n[row][col] = rand.Intn(10)
		}
	}

	fmt.Println(n)
	fmt.Println(initArray())
}

// Return a random 2D array with values 0->9
func initArray() [3][3]int {
	a := [3][3]int{}
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			a[row][col] = rand.Intn(10)
		}
	}
	return a
}
