package main

import "fmt"

func main() {
	number := 5

	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = \n", number, i)
	}
}
