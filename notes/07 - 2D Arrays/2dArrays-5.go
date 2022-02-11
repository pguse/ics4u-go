package main

import (
	"fmt"
)

func main() {
	matrix := [3][4]int{
		{2, 7, 6, 5},
		{9, 5, 1, 0},
		{4, 3, 8, 2}}

	fmt.Printf("# Rows: %d\n", len(matrix))
	fmt.Printf("# Columns: %d", len(matrix[0]))
}
