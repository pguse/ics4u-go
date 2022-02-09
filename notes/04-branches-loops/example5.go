package main

import "fmt"

func main() {
	m := 5
	n := 8

	// At least one odd - version 1
	if m%2 != 0 {
		fmt.Println("Odd")
	} else if n%2 != 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Neither one of the values is odd.")
	}

	// At least one odd - version 2
	if m%2 != 0 || n%2 != 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Neither one of the values is odd.")
	}

	// Both odd
	if m%2 != 0 && n%2 != 0 {
		fmt.Println("Both Odd")
	} else {
		fmt.Println("One of the numbers is even.")
	}
}
