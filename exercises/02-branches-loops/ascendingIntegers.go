package main

import "fmt"

func main() {
	// Declare another variable with its value
	a, b, c := 10, 15, 20

	// Are these values descending?
	if a > b && b > c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
