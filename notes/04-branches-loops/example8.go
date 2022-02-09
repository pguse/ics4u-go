package main

import "fmt"

func main() {
	n := 15
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Println(n)
	}
}
