package main

import "fmt"

func main() {
	a := 52

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%T", &a)
}
