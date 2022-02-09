package main

import "fmt"

func main() {
	marks := [3]float64{75, 80, 90} // Declare Array with 3 elements
	fmt.Printf("Your average is: %0.2f", mean(marks))
}

func mean(a [3]float64) float64 {
	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	return sum / float64(len(a))
}
