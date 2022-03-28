package main

import "fmt"

func main() {
	text := "Hello World!"

	fmt.Printf("%T\n", text)

	for i := 0; i < len(text); i++ {
		fmt.Printf("%v ", text[i])
	}

	fmt.Println()

	for i := 0; i < len(text); i++ {
		fmt.Printf("%c", text[i])
	}
}
