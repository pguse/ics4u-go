package main

import (
	"fmt"
)

func main() {
	magic := [3][3]int{
		{2, 7, 6},
		{9, 5, 1},
		{4, 3, 8}}

	fmt.Println(magic[0][:])
	fmt.Printf("%T\n", magic[0][:])
}
