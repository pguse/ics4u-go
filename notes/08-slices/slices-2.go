package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8}
	b := a
	b[2] = 10
	fmt.Println(a)
	fmt.Println(b)
}
