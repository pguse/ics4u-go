package main

import "fmt"

func main() {

	var marks = [8]int{75, 82, 90, 95, 87, 80, 70, 92}
	var average float64
	var sum int

	// Simple array output - You can comment this out
	fmt.Println(marks)

	for i := 0; i < len(marks); i++ {
		// Update the accumulator
		sum = marks[0]
	}

	fmt.Printf("Marks:  ")
	// Insert a Loop
	fmt.Printf("%d ", marks[0])

	// Correct this calculation
	average = float64(sum)
	fmt.Printf("\nAverage:  %0.1f", average)
}
