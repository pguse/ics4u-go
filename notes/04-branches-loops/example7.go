package main

import "fmt"

func main() {
	N := 5
	sum := 0
	for i := 0; i <= N; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
