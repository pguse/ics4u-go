package main

import "fmt"

func main() {
	a, b := 19, 28

	fmt.Printf("%d, %d\n", a, b)
	// output the addresses in memory (RAM) of variables a and b
	// note:  the memory address are different
	fmt.Printf("%p, %p\n", &a, &b)

	// assign variable c to the address of variable a
	// note:  c is called a pointer
	c := &a
	// assign the value 29 to the memory address stored by variable c
	*c = 29
	// note:  the value of variable a has changed
	fmt.Printf("a = %d, Type: %T\n", a, a)
	fmt.Printf("c = %p, Type: %T\n", c, c)
}
