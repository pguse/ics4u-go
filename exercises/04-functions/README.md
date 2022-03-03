# Functions

In Visual Studio Code, create a folder called **Functions** and then open it. Now open a new Terminal window. 

## Exercises

Create the following functions and implement / use them in the **main** function.

## 04-0: Slope
Write a function called **slope(x1, y1, x2, y2 float64) float64** that calculates and **returns** a **float64** value representing the slope of a line passing through the points (x1,y1) and (x2,y2).  Notice:  Since all the input parameters provided to the function are of the **same type**, we need only provide a single type declaration **float64** at the end of the list

## 04-1: Hypotenuse
Write a function called **hypotenuse(a, b float64) float64** that calculates and **returns** a **float64** value representing the hypotenuse of a right triangle with sides a and b.

## 04-2: Distance
Write a function called **distance(x1, y1, x2, y2 float64) float64** that calculates and **returns** a **float64** representing the distance between the points (x1,y1) and (x2,y2).

## 04-3: Prime
Write a function called **isPrime(num int) bool** that determines whether an integer **num** is prime and **returns** a **bool** value ... either **true** or **false** . 

## 04-4: Greatest Common Divisor
Write a function called **gcd(m, n int) int** that **returns** an **int** representing the greatest common divisor of the numbers m and n, assuming m > n. The best-known algorithm for finding a greatest common divisor is Euclid’s Algorithm. Euclid’s Algorithm states that the greatest common divisor of two integers m and n is 

* n if n divides into m completely.
* However, if n does not divide into m completely, then the answer is the **greatest common divisor** of n and the **remainder** of m divided by n.

Use the following starter code to start these exercises,

```go
package main

import "fmt"

func main() {
	fmt.Printf("Slope between (1,2) and (3,6): %0.2f\n", slope(1, 2, 3, 6))
	fmt.Printf("Hypotenuse - base: 3; height: 4: %0.2f\n", hypotenuse(3, 4))
	fmt.Printf("Distance between (1,2) and (7,10): %0.2f\n", distance(1, 2, 7, 10))
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

func isPrime(num int) bool {
	return false
}

func gcd(m, n int) int {
	return 0
}
```