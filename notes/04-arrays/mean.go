package main

import "fmt"

func main() {
	marks := [3]int{85, 95, 82} // Declare Array with 3 elements

	// Using a for-loop to add up the marks in the
	sum := 0
	for i := 0; i < len(marks); i++ {
		sum = sum + marks[i]
	}

	mean := float64(sum) / 3

	fmt.Printf("Your average is: %0.1f%%", mean)
}
