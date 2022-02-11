package main

import (
	"fmt"
)

func main() {
	m := [3][3]int{
		{1, 2, 3},
		{3, 4, 5},
		{6, 7, 8}}

	fmt.Println(m)
	fmt.Println(m[2][1])
	fmt.Println(total(m))
}

func total(a [3][3]int) int {
	s := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			s += a[row][col]
		}
	}
	return s
}
