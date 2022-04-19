package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{2, -5, 8, -10, 3, 6, 0, 15, 11}
	sort.Ints(values)
	fmt.Println(values)
}
