package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []string{"ruby", "python", "go", "c", "java", "pyret", "haskell", "racket"}
	sort.Strings(values)
	fmt.Println(values)
}
