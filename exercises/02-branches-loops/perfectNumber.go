package main

import "fmt"

func main() {
	number := 12
	sum := 0

	for i := 1; i < number; i++ {

		// Is the remainder of (number / i) equal to zero?

		sum = sum + i
	}

	if sum == number {
		fmt.Printf("The number %d is a perfect number.", number)
	}

}
