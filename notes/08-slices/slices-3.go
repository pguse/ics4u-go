package main

import "fmt"

func main() {
	a := [4]int{2, 4, 6, 8}
	b := a
	b[2] = 10
	fmt.Println(a)
	fmt.Println(b)
}
