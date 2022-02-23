package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n := [5]int{2, 4, 8, 16, 32}

	// array elements are stored in adjacent memory addresses
	fmt.Printf("%p, %p, %p, %p, %p\n", &n[0], &n[1], &n[2], &n[3], &n[4])

	// note: each int takes up 64 bits or 8 bytes of memory (RAM)
	fmt.Println(bits.UintSize)
}
