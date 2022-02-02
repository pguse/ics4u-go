package main

import "fmt"

func main() {
	// Declare another variable with its value
	a, b := 10, 15

	// Is at least one value even?
	if a%2 == 0 || b%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
