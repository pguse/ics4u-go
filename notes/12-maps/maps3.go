package main

import (
	"fmt"
	"sort"
)

func main() {
	marks := make(map[string]int) // mapping from strings to ints

	marks["alice"] = 85
	marks["charlie"] = 75
	marks["bob"] = 90

	namesSlice := make([]string, 0, len(marks))

	for k := range marks {
		namesSlice = append(namesSlice, k)
	}

	sort.Strings(namesSlice)

	for _, name := range namesSlice {
		fmt.Printf("%-10s %d\n", name, marks[name])
	}
}
