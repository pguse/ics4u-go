package main

import "fmt"

func main() {
	// Repetition
	for i := 0; i < 3; i++ {
		fmt.Println("Hello everyone!")
	}
	fmt.Println("Loop finished.")

	fmt.Println()

	// i increasing
	for i := 0; i < 3; i++ {
		fmt.Println(i, "- Hello everyone!")
	}

	fmt.Println()

	// i decreasing
	for i := 3; i > 0; i-- {
		fmt.Println(i, "- Hello everyone!")
	}
}
