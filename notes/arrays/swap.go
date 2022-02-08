package main

import "fmt"

func main() {
	marks := [3]int{75, 80, 90} // Declare Array with 3 elements

	fmt.Println(marks)

	// swap elements 1 and 2
	swap := marks[1]
	marks[1] = marks[2]
	marks[2] = swap

	fmt.Println(marks)
}
