package main

import "fmt"

func main() {
	temp := 27.0
	number := 18

	fmt.Printf("%f = %f\n", temp, *toCelsius(&temp))
	fmt.Printf("Divisors of %d: %v", number, *divisors(&number))
}

func toCelsius(fahrenheit *float64) *float64 {
	var celsius float64
	//
	// ADD CODE HERE
	//
	return &celsius
}

func divisors(n *int) *[]int {
	div := []int{}
	//
	// ADD CODE HERE
	//
	return &div
}
