# Working with Numbers

## Exercises

## 01-0: Student Average

1.  In Visual Studio Code, create a folder called **Numbers** and then open it. Now open a new Terminal window. 
2.  Create a new file called *average.go* with the following source code

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

3.  Notice the use of the Printf() function.  This allows you to create formatted output.  Change the %f to %0.2f.  What happens?
4.  Output the average with **no decimal places**.
5.  Notice how the last statement outputs the type *float64* of the m1 variable.  Change m1 from 80.0 to 80 and run the program.  Notice the error.  **Variables of different types *(int and float64)* cannot be combined.**  Change m2 to 90 and m3 to 96 and change the %f in the Printf() function to %d.  Notice that only the whole number part of the average is output - *this is a bad idea*.  Also, the type of m1 is now int.  Note:  In Printf() %f is used to output floating-point values and %d is used to output integers.

## 01-1: Area of a Circle
In Visual Studio Code, create a new file in **Numbers** called *circleArea.go*.  The user should enter the **radius** of a circle, and the program should output the circle's **area**.  **Note:** Use **Math.PI** to represent the value of pi and **Math.Pow()** to perform the exponentiation.

### For example :
```
Radius:  5
Area:  78.53981633974483
```

## 01-2: Hypotenuse of a Right-Angled Triangle
In Visual Studio Code, create a folder called **Hypotenuse** and then run the command dotnet new console in the terminal/console window.  The user should enter the base and height of a right angled-triangle, and the program should output the length of the triangle's hypotenuse.  **Note:** Use **Math.Sqrt()** to do the square root calculation.

### For example :
```
Base:  3
Height:  4
Hypotenuse: 5.0
```

## 01-3: Swap Digits *(using the modulus % operator)*
In Visual Studio Code, create a folder called **Swap** and then run the command dotnet new console in the terminal/console window.  The user should enter a **two-digit** number (example: 17), and the program should output the **number** with **swapped digits** (example: 71).  

### For example :
```
Number:  17
Swapped Digits:  71
