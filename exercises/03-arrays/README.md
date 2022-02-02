# Arrays

## Exercises

Create a folder called **Arrays**.

## 03-0: Student Average

1.  Create a file in **Arrays** called **average.go**.  Modify the starter code below so that it averages 8 marks given in an integer array.  You will calculate the average, by using a **for-loop**.  ***Note:  You will need to use the idea of using an accumulator that we used in grade 11 with Python.***
2.  Output both the marks ***(all on one line)*** and the average, as shown in the example below.

```go
package main

import "fmt"

func main() {

	var marks = [8]int{75, 82, 90, 95, 87, 80, 70, 92}
	var average float64
	var sum int

	// Simple array output - You can comment this out
	fmt.Println(marks)

	for i := 0; i < len(marks); i++ {
		// Update the accumulator
		sum = marks[0]
	}

	fmt.Printf("Marks:  ")
	// Insert a Loop
	fmt.Printf("%d ", marks[0])

	// Correct this calculation
	average = float64(sum)
	fmt.Printf("\nAverage:  %0.1f", average)
}
```


**Example Output:**
```
Marks:  72 85 67 78 90 81 87 75
Average:  79.4
```

## 03-1: Reverse Output

Create a file in **Arrays** called **reverseOutput.go**.  Modify the starter code below so that it displays the given array in reverse order.

```go
package main

import "fmt"

func main() {

	var values = [3]int{2, 5, 7}

	// Insert a loop
	fmt.Printf("%d ", values[0])
}
```

For a given array: [2 5 7]

**Example Output**
```
7 5 2
```

## 03-2: Copying Elements Array-to-Array in Reverse Order

Create a file in **Arrays** called **arrayCopyReverse.go**.  Modify the source code below to copy the elements in one array into another array, but in reverse order.  

```go
package main

import "fmt"

func main() {

	var values = [4]int{2, 8, 5, 10}
	var reverseValues [4]int

	for i := 0; i < len(values); i++ {
		reverseValues[i] = values[i]
	}

	// Insert Loop
	fmt.Printf("%d ", reverseValues[0])
}
```

For the given array: [2, 8, 5, 10]

**Example Output**
```
10 5 8 2
```

## 03-03: Maximum and Minimum Elements in an Array

Create a file in **Arrays** called **maxMin.go*.  Modify the starter code below to find the maximum and minimum element in a given array of positive integers.  

```go
package main

import "fmt"

func main() {

	var values = [5]int{40, 25, 15, 70, 60}
	var max int
	var min int

	for i := 0; i < len(values); i++ {
		// Insert an if-statement
		max = values[i]
		// Insert an if-statement
		min = values[i]
	}

	fmt.Printf("Maximum:  %d\n", max)
	fmt.Printf("Minimum:  %d", min)
}
```

For a given array: [40, 25, 15, 70, 60]

**Example Output**
```
Maximum: 70
Minimum: 15
```
