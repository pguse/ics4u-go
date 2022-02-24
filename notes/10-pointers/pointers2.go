package main

import "fmt"

func main() {

	value := 3

	fmt.Println(addTwo(value))
	fmt.Printf("value: %v\n", value)

	addTwoPointer(&value)
	fmt.Printf("value: %v\n", value)
}

func addTwo(n int) int {
	return n + 2
}

func addTwoPointer(n *int) {
	*n += 2
}
