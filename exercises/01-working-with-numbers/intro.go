package main

import "fmt"

func main() {
	// Two ways to declare a variable in Go

	//var name string = "Mr. Guse"
	name := "Mr. Guse"

	// Two way to output a message
	//fmt.Println("Hello", name)
	fmt.Printf("Hello %v", name)
}
