package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []int{5, 2, 1, 3, 6, 8, -5}
	fmt.Println(a)
	fmt.Println(selSort(a))
}

// Sorts an int slice using the selection sort algorithm
func selSort(values []int) []int {
	v := make([]int, len(values))
	copy(v, values)

	// 	ADD CODE HERE

	return v
}

// Sorts an int slice using the bubble sort algorithm
func bubSort(values []int) []int {
	v := make([]int, len(values))
	copy(v, values)

	// 	ADD CODE HERE

	return v
}

// Shuffles an int slice
func shuffle(values []int) []int {
	m := make([]int, len(values))
	copy(m, values)
	for i := range m {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(len(m))
		swap := m[i]
		m[i] = m[n]
		m[n] = swap
	}

	return m
}
