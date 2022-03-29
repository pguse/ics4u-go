package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func main() {

	var p1 point
	p1.x = 4
	p1.y = 5

	p2 := point{1, 5}
	p3 := point{x: 7, y: 9}

	fmt.Printf("Point #1: (%0.1f, %0.1f)\n", p1.x, p1.y)

	fmt.Printf("Points: %v %v\n", p1, p2)
	fmt.Printf("Midpoint: %v\n", midpoint(p1, p2))
	fmt.Printf("Distance: %0.2f", distance(p1, p3))
}

func midpoint(p1, p2 point) point {
	return point{(p1.x + p2.x) / 2, (p1.y + p2.y) / 2}
}

func distance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
