package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 5.0

	area := math.Pi * math.Pow(radius, 2)

	fmt.Printf("Area of the Circle: %f\n", area)
}
