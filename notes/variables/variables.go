package main

import "fmt"

func main() {
	name := "Albert"
	age := 17
	mark := 82.5

	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of mark: %T\n", mark)

	fmt.Printf("Value of name: %s\n", name)
	fmt.Printf("Value of age: %d\n", age)
	fmt.Printf("Value of mark: %f\n", mark)
}
