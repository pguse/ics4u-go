package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//a := []int{2, 4, 6, 8, 10, 12, 14}
	b := []int{5, 2, 1, 3, 6, 8, -5}
	//c := []int{2, 8, 43, 9, 22, 12, 14}
	fmt.Println(b)
	fmt.Println(selSort(b))
}

func selSort(values []int) []int {
	v := make([]int, len(values))
	copy(v, values)

	// Search for the minimum value starting at 0
	minIndex := 0
	for i := 0; i < len(v); i++ {
		if v[i] < v[minIndex] {
			minIndex = i
		}
	}
	// Swap the values
	swap := v[0]
	v[0] = v[minIndex]
	v[minIndex] = swap

	// Search for the minimum value starting at 1
	minIndex = 1
	for i := 1; i < len(v); i++ {
		if v[i] < v[minIndex] {
			minIndex = i
		}
	}
	// Swap the values
	swap = v[1]
	v[1] = v[minIndex]
	v[minIndex] = swap

	return v
}

func bubSort(values []int) []int {
	return []int{}
}
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
