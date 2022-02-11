package main

import "fmt"

func main() {

	v := [5]int{4, 5, -3, 0, 7}

	w := v[:4]

	fmt.Println(w)
}
