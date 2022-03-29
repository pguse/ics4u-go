package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 65
	//y := string(x)
	y := strconv.Itoa(x)
	fmt.Println(x, y)
}
