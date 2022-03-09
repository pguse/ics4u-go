package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func main() {

	p1 := Point{4, 5}
	p2 := Point{7, 9}

	fmt.Printf("Points: %v %v\n", p1, p2)
	fmt.Printf("Midpoint: %v\n", midpoint(p1, p2))
	fmt.Printf("Distance: %0.2f", distance(p1, p2))
}

func midpoint(p1, p2 Point) Point {
	return Point{(p1.x + p2.x) / 2, (p1.y + p2.y) / 2}
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
