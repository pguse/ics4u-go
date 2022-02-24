package main

import "fmt"

func main() {
	numberSlice := []int{87, 65, 92, 99}

	fmt.Println("Doubling a Slice")
	fmt.Println(numberSlice)
	fmt.Println(doubleSlice(numberSlice))
	fmt.Println(numberSlice)
}

func doubleSlice(n []int) []int {
	m := make([]int, len(n))
	copy(m, n)
	for i := 0; i < len(m); i++ {
		m[i] *= 2
	}
	return m
}
