# Slices

In Go, arrays are actually not used very often for a couple of reasons.

1. An array has a fixed size chosen when the array is first created. You cannot remove or add an element to an array, except by creating another array with this new size.

 2. Since **its size is part of its type**, using an array as an input parameter or return value of a function is very restrictive.  A function can only be defined for an array of a single size, so you would need different functions for different sized arrays.

 In Go, a **slice** is a type that allows storage of a sequence of values with a variable length.  Here is an example of a slice being declared and initialized,

 ```go
 numberSlice := []int{87, 65, 92}
 ```

In order to add an element to a slice, you use the **append** function.  For example,

```go
numberSlice = append(numberSlice, 99)
fmt.Println(numberSlice)
```

produces the output

```
[87 65 92 99]
```

## Concatenating Two Slices

Let's say we want to concatenate two slices.  If we try to use the **append** function, the following code

```go
numberSlice := []int{87, 65, 92}
anotherSlice := []int{15, 76, 44}
result := append(numberSlice, anotherSlice)
```

produces the error

```
cannot use anotherSlice (type []int) as type int in append
```

because the **append** function is expecting a **slice** and a single value as arguments.  However, **append** is called a **variadic** function, meaning it actually can accept multiple arguments.  For example, the following code,

```go
numberSlice := []int{87, 65, 92}
numberSlice = append(numberSlice, 54, 35)
fmt.Println(numberSlice)
```

will work, producing the output

```
[87 65 92 54 35]
```

In order to **append** a slice to another slice, you must use **...** at the end of the slice name.  The following code

```go
numberSlice := []int{87, 65, 92}
anotherSlice := []int{15, 76, 44}
result := append(numberSlice, anotherSlice...)
fmt.Println(result)
```

produces the output

```
[87 65 92 15 76 44]
```

without error, because the **...** converts the slice to a group of multiple arguments, which the **append** function can handle.

## Using Slices in Functions

Passing a slice into a function allows you to create a single function that acts on variable length sequences of values.  The following function,

```go
func average(n []int) float64 {
	s := 0
	for i := 0; i < len(n); i++ {
		s = s + n[i]
	}

	return float64(s) / float64(len(n))
}
```

can be used with a slice of integers of any length.  For example,

```go
fmt.Printf("Average: %0.2f\n", average(numberSlice))
```

produces the output,

```
Average: 85.75
```

and would work for a slice of any _non-zero_ number of integers.

## A Contrast - Slices and Arrays

Let's compare two similar functions.  This one acts on a slice,

```go
func doubleSlice(n []int) {
	for i := 0; i < len(n); i++ {
		n[i] *= 2
	}
}
```

Notice that even through this function does not have a return value. The following code,

```go
fmt.Println(numberSlice)
doubleSlice(numberSlice)
fmt.Println(numberSlice)
```

produces the output

```
[87 65 92 99]
[174 130 184 198]
```

Notice how the original slice was modified, even though no value was returned from the function.  This is in contrast to what happens when you use an array.  For example, the following function

```go
func doubleArray(n [4]int) {
	for i := 0; i < len(n); i++ {
		n[i] *= 2
	}
}
```

doesn't have much of a purpose, as the following code,

```go
numberArray := [4]int{54, 78, 23, 87}
fmt.Println(numberArray)
doubleArray(numberArray)
fmt.Println(numberArray)
```

produces the following output,

```
[54 78 23 87]
[54 78 23 87]
```

Go is a **copy by value** programming language.  When you pass a value into a function a copy is always made, which means that modifying the copy (in this case **n**) in general does not modify the original value.  However, when slices are copied, the copy is always **linked** to the original.  For example,

```go
a := []int{2, 4, 6, 8}
b := a
b[2] = 10
fmt.Println(a)
fmt.Println(b)
```

produces the output

```
[2 4 10 8]
[2 4 10 8]
```

Notice how changing slice **b** also changes slice **a**.  However, if you run the same code using arrays,

```go
a := [4]int{2, 4, 6, 8}
b := a
b[2] = 10
fmt.Println(a)
fmt.Println(b)
```

you get the output

```
[2 4 6 8]
[2 4 10 8]
```

that you probably expect because arrays are **copied by value**.  A copy of an array has no connection to the original array, unlike slices.

## Functions with Immutable Slices

You can create a copy of a slice that has **no connection** with the original.  To do this, you must use the **copy** function.  For example,

```go
a := []int{2, 4, 6, 8}
b := make([]int, 4)
copy(b, a)
b[2] = 10
fmt.Println(a)
fmt.Println(b)
```

produces the output

```
[2 4 6 8]
[2 4 10 8]
```

Notice how the original slice **a** has not been affected by changing the copy **b**.  The **copy** function copies a slice **a** to another slice **b** based on the number of elements in **b**.  To copy all the elements of **a**, **b** must have the same length.  The **make** function creates an integer **slice** with a length of 4, containing the default values for the type _(0 for **int**)_.

Now, back to our **doubleSlice** function.  The following version,

```go
func doubleSlice(n []int) []int {
	m := make([]int, len(n))
	copy(m, n)
	for i := 0; i < len(m); i++ {
		m[i] *= 2
	}
	return m
}
```

makes a copy **m** of the slice copied into the function using the **copy** function.  Modifying **m** now **does not modify** the original slice as the following code,

```go
numberSlice := []int{87, 65, 92, 99}

fmt.Println("Doubling a Slice")
fmt.Println(numberSlice)
fmt.Println(doubleSlice(numberSlice))
fmt.Println(numberSlice)
```

produces the output

```
[87 65 92 99]
[174 130 184 198]
[87 65 92 99]
```

Be careful when passing slices into functions.  Make sure that you understand whether you intend on modifying the original slice or whether you would prefer to keep the original slice **immutable**.