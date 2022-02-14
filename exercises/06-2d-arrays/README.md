# Functions

In Visual Studio Code, create a folder called **2D Arrays** and then open it. Now open a new Terminal window. 

## Exercises

## 1 - Random Matrix

Create a function called **randMatrix** that creates a 2-dimensional array, with **3** rows and **3** columns, made up of 0's and 1's chosen randomly.  The function should return the array that gets created.

```go
func randMatrix() [3][3]int {
}
```

## 2 - Find

Create a function called **find** that searches for an integer **x** in a 3x3 integer array **m**.  It should return the position **row, col** if it finds it, otherwise it returns **-1,-1**.

```go
func find(m [3][3]int, x int) int, int {
}
```

## 3 - Maximum

Create a function called **max** that returns the maximum value in a 3x3  integer array **m**.

```go
func max(m [3][3]int) int {
}
```

## 4 - Addition

Create a function called **add** that adds two 3x3 integer **2D arrays** and returns a new 2-dimensional array. **Addition of arrays** produces a new array of the same size by adding **elements** from the original arrays with the **same row and column**, and placing the **sum** in the **same position** in the new array.

```go
func add(m [3][3]int, n [3][3]int) [3][3]int {
}
```

## 5 - Scalar

Create a function called **scalar** that multiplies every value of a 3x3 2D array **m** by the value **k** and returns the new 2-dimensional array."

```go
func scalar(m [3][3]int, k int) [3][3]int {
}
```

## 6 - Sum of Neighbours

Create a function called **sumNeighbours** that finds the sum of the neighbours of any value in the position _(row,col)_ of a 3x3 2-dimensional array.  A neighbour is any value left, right, above, below, or diagonal to the position _(row,col)_.  The function should return the number of neighbours."

```go
func sumNeighbours(m [3][3], row int, col int) int {
}
```
