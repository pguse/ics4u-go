package main

import "fmt"

func main() {
	var numbers = [4]int{6, 28, 496, 8128}
	var cheeses = [5]string{"Cheddar", "Edam", "Gouda", "Havarti", "Mozzarella"}

	fmt.Printf("%T\n", numbers)
	fmt.Printf("%T", cheeses)
}
