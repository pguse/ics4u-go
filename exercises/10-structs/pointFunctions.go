package main

import (
	"fmt"
)

type Point struct {
	x float64
	y float64
}

func main() {

	p1 := Point{4, 5}
	p2 := Point{x: 7, y: 9}

	fmt.Printf("Distance: %0.2f", distance(p1, p2))
	fmt.Printf("Slope: %0.2f", slope(p1, p2))

}

func distance(p1, p2 Point) float64 {
	return 0.0
}

func slope(p1, p2 Point) float64 {
	return 0.0
}
