package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{2, 4, 6, 8, 10, 12, 14}

	fmt.Println(a)
	fmt.Println(shuffle(a))
	fmt.Println(a)
}

func shuffle(values []int) []int {
	m := make([]int, len(values))
	copy(m, values)
	for i := range m {
		n := rand.Intn(len(m))
		swap := m[i]
		m[i] = m[n]
		m[n] = swap
	}

	return m
}
