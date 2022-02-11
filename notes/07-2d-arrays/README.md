# 2-Dimensional Arrays

A **1 dimensional array** is a **linear** representation of elements.  It represents a type of list that can be accessed by subsequent memory locations. For example:

```go
v := [5]int[4, 5, -3, 0, 7]
```
where each element of the array can be accessed by an ordered **index** starting from 0.  For example:

```go
v[3]
```

A **2 dimensional array** is a representation of items in the form of **rows** and **columns**.  This is called a **tabular** representation of data.  Elements in a matrix 2D array are accessed by using a **row** and **column** index.  Another way to think of a 2D array, is an array of arrays.  For example,

```go
m := [3][3]int{{1, 2, 3}, {3, 4, 5}, {6, 7, 8}}

fmt.Println(m)
```

The output for the above code is

```
[[1 2 3] [3 4 5] [6 7 8]]
```

You can also create a **2D array** by writing the **rows** on **separate lines** as follows:

```go
m := [3][3]int{
	{1, 2, 3},
	{3, 4, 5},
	{6, 7, 8}}

fmt.Println(m)
```

The output is unchanged.

## Accessing an Element
An element in a **2D array** is accessed by using the **row** and **column** values as follows:

```go
fmt.Prinln([2][1])
```

produces the output

```
7
```

## Using a For Loop
A for-loop can be used to scan through the rows and columns of a 2 dimensional array or generate the values of the elements.

### Adding the elements

The following code adds up the elements of a 3x3 array.

```go
func total(a [3][3]int) int {
	s := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			s += a[row][col]
		}
	}
	return s
}
```

### Generating elements

You can generate a **random** array of elements by using a for-loop and the **math/rand** package.  The following code creates a 3x3 array of random values from 1->9 inclusive, in two different ways; one way uses a function.

```go
import (
	"fmt"
	"math/rand"
)

func main() {
	n := [3][3]int{}

	// Create a random 2D array with values 0->9
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			n[row][col] = rand.Intn(10)
		}
	}

	fmt.Println(n)
	fmt.Println(initArray())
}

// Return a random 2D array with values 0->9
func initArray() [3][3]int {
	a := [3][3]int{}
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			a[row][col] = rand.Intn(10)
		}
	}
	return a
}
```

The Intn() function returns, as an **int**, a non-negative pseudo-random number in the half-open interval [0,n).

## Slices

A slice of any array can be accessed using the following notation

```go
v := [5]int{4, 5, -3, 0, 7}
w := v[2:4]
fmt.Println(w)
```

The output of this code is

```
[-3 0]
```

The identifier **w** is a **slice** of array **v** from indices 2 **up to**, but not including, 4.  If you would like to create a slice up to/from one of the ends of the array, just remove one of the index values.  For example, 

```go
v := [5]int{4, 5, -3, 0, 7}
w := v[2:]
fmt.Println(w)
```

produces the output

```
[-3 0 7]
``

and

```go
v := [5]int{4, 5, -3, 0, 7}
w := v[:4]
fmt.Println(w)
```

produces the output

```
[4 5 -3 0]
```

## Slices of 2-Dimensional Arrays

This notation can also be used to access any number of rows or columns of a 2-dimensional array.  For example, we can access a single **row** of a 2D array as follows,

```go
magic := [3][3]int{
		{2, 7, 6},
		{9, 5, 1},
		{4, 3, 8}}

fmt.Println(magic[0][:])
fmt.Printf("%T", magic[0][:])
```

This produces the output

```
[2 7 6]
[]int
```

The expression ```magic[0][:]``` creates a **slice** of the magic array consisting of **row 0** and all of its elements.  **Note:**  Using ```[:]``` creates a slice of an array from its beginning to its end _(both index values are missing)_.

Notice that ```magic[0][:]``` has the type ```[]int```, which is a **slice** of integers.  If it was an array, its **size** would be included in the type _(see our note on Arrays)_.

**Note:**  We cannot access the columns of a 2D array using slices because the data is stored as an **array of arrays**, where each array represents a row.

## Dimensions of a 2-Dimensional Array

Assuming we are dealing with a rectangular 2D array, we can calculate the dimensions as follows,

```go
matrix := [3][4]int{
	{2, 7, 6, 5},
	{9, 5, 1, 0},
	{4, 3, 8, 2}}

fmt.Printf("# Rows: %d\n", len(matrix))
fmt.Printf("# Columns: %d", len(matrix[0]))
```

This produces the output

```
# Rows: 3
# Columns: 4
```