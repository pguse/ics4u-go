# Functions, Pointers, and Slices

The **Go** programming language uses a _copy-by-value_ model when using functions.  Below, we will look at some examples of how this model works.

## [Example](https://go.dev/play/p/hkOU3FASKvn) Using Values of Type **int**

_**Note:**  You can try the code described below at the link above._

Here is an example of a function that has a single **int** value as a parameter,

```go
func doubleInt(n int) {
	n *= 2
}
```

When it is used as follows,

```go
a := 5
doubleInt(a)
fmt.Println(a)
```

The **int** value stored as variable **a** is a copy of the variable **n** in the **doubleInt** function.  As a result, when the value of **n** is doubled inside the function, the value of **a** remains unchanged.  We see this in the output of this code,

```
5
```

The value of **n** inside the function only exists within the function _(called a **local** variable)_ and is not connected to the variable **a**.

In contrast, the function **doubleIntReturn**,

```go
func doubleIntReturn(n int) int {
	n *= 2
	return n
}
```

returns the value of **n** from the function,

```go
a := 5
a = doubleIntReturn(a)
fmt.Println(a)
```

which assigns that value as the new value of **a**, as shown above.  Now the output is

```
10
```

If you choose, you can use a **pointer** instead of using a **return** statement.  For example, the **doubleIntPointer** function,

```go
func doubleIntPointer(n *int) {
	*n *= 2
}
```

uses an **int** pointer as a parameter.  So, **n** stores a memory address in **RAM**.  To access the **value** stored at this memory address, you use the expression ***n**.  So, the statement ```*n *= 2``` doubles the value stored at the memory address **n**.  When this function is called, as follows

```go
a := 5
doubleIntPointer(&a)
fmt.Println(a)
```

the memory address of variable **a** is passed into the function as an argument.  This memory address, stored as **&a** is still copied to the pointer **n** in the function **doubleIntPointer**.  The value stored at that memory address is doubled inside the function, but that also changes the value stored by **a** because it is stored at the **same memory address**.  So, the output is

```
10
```

## [Example](https://go.dev/play/p/3WpY-mXo9qA) Using a Slice of Values of Type **int**

_**Note:**  You can try the code described below at the link above._

The exception to the **copy-by-value** model in **Go** happens when using **slices**.  **Slices** are **copied-by-reference** when they are a parameter in a function.  Behind the scenes, it is the **memory address** of the first element of the slice that is copied when the function is called.  So, a slice acts like a **pointer**.  For example, the function **doubleSlice**

```go
func doubleSlice(n []int) {
	for i := range n {
		n[i] *= 2
	}
}
```

will double all values of the slice that is passed in as an argument to **doubleSlice**.  So, the code

```go
num := []int{1, 2}
doubleSlice(num)
fmt.Println(num)
```

produces the output

```
[2 4]
```

Modifying the slice **n** within the function **doubleSlice** also modifies the **num** slice that was passed into the function.  This is in contrast to what we showed above for **int** values.

### Be Careful

You do need to be aware of where this behaviour of slices does not seem to behave in the manner described above.  For example, when the function

```go
func appendThree(n []int) {
	n = append(n, 3)
}
```

is called below,

```go
num := []int{1, 2}
appendThree(num)
fmt.Println(num)
```

the following output is produced

```
[1 2]
```

The reason that the original slice **num** is not modified by the **appendThree** function is that the **append** function makes a copy of **n** in the statement,

```go
n = append(n, 3)
```

This new copy of **n** is stored in a new **memory address** that is different than the slice **num** that was used as the argument of the function **appendThree**.  If you want to modify the original slice **num** when using **append** within a function, you must use a **pointer** as follows,

```go
func appendThreePointer(n *[]int) {
	*n = append(*n, 3)
}
```

This function uses a **pointer** to a **slice of integers** as the parameter.  It is called like this,

```go
num := []int{1, 2}
appendThreePointer(&num)
fmt.Println(num)
```

Within the function **appendThreePointer**, the variable **n** has the same memory address as **num**.  In the statement,

```go
*n = append(*n, 3)
```

the **append** function places the new slice in the same memory address as **n**, which is the same memory location as the slice **num**.  So, the value of **num** is modified and the output is,

[1 2 3]