package main

import "fmt"

func main() {
	marks := make(map[string]int) // mapping from strings to ints

	marks["alice"] = 85
	marks["charlie"] = 75
	marks["bob"] = 90

	for k := range marks {
		fmt.Printf("%-10s %d\n", k, marks[k])
	}

	fmt.Println(len(marks))
}
