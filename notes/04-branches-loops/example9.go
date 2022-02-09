package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		n := rand.Intn(5)

		fmt.Println(n)

		if n == 0 {
			break
		}
	}
}
