# Working with Numbers

In Visual Studio Code, create a folder called **Numbers** and then open it. Now open a new Terminal window. 

## Exercises

## 01-0: Student Average

Create a new file called *average.go* with the following source code

```go
package main

import "fmt"

func main() {
	m1 := 80.0
	m2 := 90.0
	m3 := 96.0

	average := (m1 + m2 + m3) / 3

	fmt.Printf("Average: %f\n", average)
	fmt.Printf("Type of variable m1: %T", m1)
}
```

3.  Notice the use of the **Printf()** function.  This allows you to create formatted output.  Change the %f to %0.2f.  What happens?  Notice the *\n* inside the quote.  This creates a **new line** after **Prinf()** executes.  It does not do this on its own like the **Println()** function.
4.  Output the average with **no decimal places**.
5.  Notice how the last statement outputs the type *float64* of the m1 variable, by using the *%T* character.  Change m1 from 80.0 to 80 and run the program.  Notice the error.  **Variables of different types *(int and float64)* cannot be combined.**  Change m2 to 90 and m3 to 96 and change the %f in the Printf() function to %d.  Notice that only the whole number part of the average is output - *this is a bad idea* for accuracy reasons.  Also, the type of variable m1 is now **int**.  Note:  In **Printf()** the character **%f** is used to output floating-point values and character **%d** is used to output integers.

## 01-1: Volume of a Cylinder
In Visual Studio Code, create a new file in **Numbers** called *cylinderVolume.go*.  Use the source code provided below as starter code in your file.  Modify the code so that it calculates the volume of a cylinder <img src="https://render.githubusercontent.com/render/math?math=V = \pi^2 rh"> instead of the area of a circle.  **Note:** Pay attention to the use of **math.Pi** to represent the value of pi and **math.Pow()** to perform the exponentiation.  See how the **math** package was imported by using parentheses around the two imported packages.

```go
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
```

## 01-2: Hypotenuse of a Right-Angled Triangle
In Visual Studio Code, create a new file in **Numbers** called *hypotenuse.go*.  Modify the starter code below so that it calculates the correct hypotenuse of a triangle with sides **a** and **b**.  **Note:** Use **math.Sqrt()** to do the square root calculation.

```go
package main

import (
	"fmt"
)

func main() {
	a := 3.0
	b := 4.0

	hypotenuse := a + b

	fmt.Printf("Hypotenuse: %f\n", hypotenuse)
}
```

## 01-3: Swap Digits *(using the modulus % operator)*
In Visual Studio Code, create a new file in **Numbers** called *swap.go*.  Modify the starter code below so that it swaps the digits of the variable *number*. You should save the new value in a variable called *swap*.  *Assume that the number you are swapping only has two digits.*  **Note:** The */* operator returns the **quotient** of a division of two integers, while the **modulus** operator *%* returns the **remainder** of a division of two integers.

```go
package main

import (
	"fmt"
)

func main() {
	number := 75

	fmt.Printf("Tens: %v\n", number/10)
	fmt.Printf("Remainder: %v\n", number%10)
}
```