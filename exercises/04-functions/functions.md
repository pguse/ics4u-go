# Functions

Our focus in this course is to create modular code.  One way we can accomplish this is to organize our code into functions.  Here is an example ***(see the given file - functions.go)*** of using a function in our student information program:

```go
package main

import "fmt"

func main(){
    marks := [3]int{75,80,90}; // Declare Array with 3 elements
    fmt.Prinf("Your average is: %0.2f", mean(marks))
}

func mean(a [3]int) float64 {
    sum := 0
    for i := 0; i < len(a); i++)
    {
        sum = sum + a[i]
    }
    return float64(sum) / len(a)
}
```
Notice the creation of the function:  **mean**.  It has a few characteristics that need to be pointed out.  Notice that in this example there is one input **parameter** required ... an **array of integers of size 3** called **a** _(declared within the parentheses)_.  Notice the **return type** ... in this case **float64** _(declared after the parentheses)_.  The value of the calculation is returned from the function using the keyword **return**.  If you look in the **main** function, the **mean** function acts like a variable except that an array called marks is passed into it.  Its return value is output using the **Printf** function.

Functions can exist that do not return values.  An example is the **main** function we have been using in all our programs.  It has neither any input **parameters** provided or values returned.

## Exercises

Create the following functions and implement / use them in the **main** function.

## 04-1: Slope
Write a function called **slope(x1, y1, x2, y2 float64) float64** that calculates and **returns** a **float64** value representing the slope of a line passing through the points (x1,y1) and (x2,y2).  Notice:  Since all the input parameters provided to the function are of the **same type**, we need only provide a single type declaration **float64** at the end of the list

## 04-2: Hypotenuse
Write a function called **hypotenuse(a, b float64) float64** that calculates and **returns** a **float64** value representing the hypotenuse of a right triangle with sides a and b.

## 04-3: Distance
Write a function called **distance(x1, y1, x2, y2 float64) float64** that calculates and **returns** a **float64** representing the distance between the points (x1,y1) and (x2,y2).

## 04-4: Prime
Write a function called **isPrime(num uint) bool** that determines whether an unsigned integer **num** is prime and **returns** a **bool** value ... either **true** or **false** . 

## 04-5: Greatest Common Divisor
Write a function called **gcd(m, n uint) uint** that **returns** an **uint** representing the greatest common divisor of the numbers m and n, assuming m > n. The best-known algorithm for finding a greatest common divisor is Euclid’s Algorithm. Euclid’s Algorithm states that the greatest common divisor of two integers m and n is 

* n if n divides m evenly.
* However, if n does not divide m evenly, then the answer is the **greatest common divisor** of n and the **remainder** of m divided by n.

Use the following starter code to start these exercises,

```go
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
```