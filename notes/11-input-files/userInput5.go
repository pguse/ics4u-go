package main

import "fmt"

func main() {

	var first string
	var last string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&first, &last)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Printf("Hello %s %s!\n", first, last)
	fmt.Printf("Your are %d years old.", age)
}
