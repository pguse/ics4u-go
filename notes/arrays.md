# Arrays
In computer program, we sometimes want to store groups of related items. In Go, we can use arrays for this purpose.

An array is an ordered sequence of values.  The values in an array are called _elements_ or sometimes _items_. There are several ways to create a new array; the simplest is to enclose elements in curly braces {}

## Creating Simple Arrays

Here is an example of a simple one-dimensional array.

```go
numbers := [4]int{6, 28, 496, 8128}
```

The function ```fmt.Printf()``` can be used to find out the type of the array."

```go
fmt.Printf("%T", numbers)
```
**Output:**
```
[4]int
```

The type of array is composed of a number, representing its size, and a data type.  The number indicates the size. In our example the elements are of type **int** and the array has 4 elements.  Unlike other pieces of data, such as integers and floats, the size of an array is inherently part of its type.  Here is another example of an array.

```go
cheeses := [5]string{"Cheddar", "Edam", "Gouda", "Havarti", "Mozzarella"}

fmt.Printf("%T", cheeses)
```

**Output:**
```
[5]string
```

This array has size 5 and contains values of type **string**."

## Arrays are Mutable

The syntax for accessing the elements of an array uses the **bracket operator []**.  Like a number of computer languages, the indices of an array in Go start at 0.

```go
fmt.Println(cheeses[0])
```

**Output:**
```
Cheddar
```

The phrase **arrays are mutable** means that we are able to change an element of an array.  For example,

```go
numbers[3] = 89
fmt.Println(numbers)
```

Now the **numbers** array looks like this.

```
[6 28 496 89]
```

## Length of an Array

Go provides the **len()** function that returns the length of an array."

```go
fmt.Println(len(cheeses))
```

**Output**
```
5
```

When copying an array to another array in other languages like c/c++/python, a **reference** to the array is copied, and any changes made in the copied array will change the original array.  In Go this is called a **pointer**.  We will discuss **pointers** in a later note.

However, in Go when an array is copied to another variable a new independent copy is created.  Making a change to either array does not affect the other one.  For example,

```go
values := numbers
numbers[2] = 19
fmt.Println(numbers)
fmt.Println(values)
```

**Output**
```
[6 28 19 89]
[6 28 496 89]
```

## For loops - Traversing an Array

A **for loop** can be used to to traverse the elements of an array. For example, the following code moves through the elements of an array of marks and adds them to an existing **sum**.  

```go
package main

import "fmt"

func main() {
	marks := [3]int{85, 95, 82} // Declare Array with 3 elements

	// Using a for-loop to add up the marks in the array
	sum := 0

	for i := 0; i < len(marks); i++ {
		sum = sum + marks[i]
	}

	mean := float64(sum) / 3

	fmt.Printf("Your average is: %0.1f%%", mean)
}
```

## Swapping Elements

Sometimes you need to rearrange elements in an array. You do this by swapping elements as in the following example code:

```go
package main

import "fmt"

func main(){
  marks := [3]int{75, 80, 90} // Declare Array with 3 elements
  
  fmt.Println(marks)

  // swap elements 1 and 2
  swap := marks[1]
  marks[1] = marks[2]
  marks[2] = swap
  
  fmt.Println(marks)
}
```

**Output:**
```
[75 90 80]
```