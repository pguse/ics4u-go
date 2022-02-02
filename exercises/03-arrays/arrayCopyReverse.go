package main

import "fmt"

func main() {

	var values = [4]int{2, 8, 5, 10}
	var reverseValues [4]int

	for i := 0; i < len(values); i++ {
		reverseValues[i] = values[i]
	}

	// Insert Loop
	fmt.Printf("%d ", reverseValues[0])
}
