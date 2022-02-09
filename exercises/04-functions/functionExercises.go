package main

import "fmt"

func main() {
	fmt.Printf("Slope between (1,2) and (3,6): %0.2f\n", slope(1, 2, 3, 6))
	fmt.Printf("Hypotenuse - base: 3; height: 4: %0.2f\n", hypotenuse(3, 4))
	fmt.Printf("Distance between (1,2) and (7,10): %0.2f\n", distance(1, 2, 7, 10))
	fmt.Printf("Slope between (1,2) and (3,6): %0.2f\n", slope(1, 2, 3, 6))
	fmt.Printf("Is 15 prime? %t\n", isPrime(15))
	fmt.Printf("GCD of 48 and 36: %d\n", gcd(48, 36))
}

func slope(x1, y1, x2, y2 float64) float64 {
	return 0.0
}

func hypotenuse(a, b float64) float64 {
	return 0.0
}

func distance(x1, y1, x2, y2 float64) float64 {
	return 0.0
}

func isPrime(num uint) bool {
	return false
}

func gcd(m, n uint) uint {
	return 0
}
