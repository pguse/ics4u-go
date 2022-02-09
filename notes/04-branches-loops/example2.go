package main

import "fmt"

func main() {
	mark := 72

	if mark >= 80 {
		fmt.Println("A")
	} else if mark >= 70 {
		fmt.Println("B")
	} else if mark >= 60 {
		fmt.Println("C")
	} else if mark >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("You are not meeting course expectations.")
	}
}
