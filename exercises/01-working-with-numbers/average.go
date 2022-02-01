package main

import "fmt"

func main() {
	m1 := 80.0
	m2 := 90.0
	m3 := 96.0

	average := (m1 + m2 + m3) / 3

	fmt.Printf("Average: %f\n", average)
	fmt.Printf("Type of variable m1: %T", m1)
}
