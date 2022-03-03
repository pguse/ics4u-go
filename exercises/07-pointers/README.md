# Pointers

In Visual Studio Code, create a folder called **Pointers** and then open it. Now open a new Terminal window. 

## Exercises

Complete the following exercises.  Starter code is found in the **exercises.go** file.  To see an example of how to use a **pointer to a slice**, look at the **pointers.go** file.

## 0 - Celsius

Write a function called **toCelsius(fahrenheit \*float64) \*float64** that takes a **pointer** to a temperature value in **fahrenheit** as a parameter.  It should return a **pointer** to the equivalent temperature value in **celsius**.  The conversion equation is:

<img src="https://render.githubusercontent.com/render/math?math=C = \frac{5}{9}(F-32)} = -1">

## 1 - Divisors

Write a function called **divisors(n \*int) \*[]int** that takes a **pointer** to an integer value and **returns** a **pointer** to a slice of integers containing the divisors.