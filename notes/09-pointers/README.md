# Pointers

## Variables

Variables in Go have four properties:  a name, a type, a value, and a memory address _(where the value is stored)_.  You can access the **value** as follows,

```go
a := 52
fmt.Println(a)
```

This outputs

```
52
```

You can access the type as follows,

```go
fmt.Printf("%T", a)
```

This outputs

```
int
```

To access the memory address, the following code

```go
fmt.Printf("%p", &a)
```

outputs, on my machine, 

```
0xc0000aa058
```

If you run this on your machine, the memory address will be different. Notice the use of the **&** character in front of the variable **a**.

## Pointer Variables

A pointer is a variable that stores the value of a memory address in RAM.  To declare a pointer, you can do the following,

```go
var c *int
```

The **\*int** indicates that c is a pointer to an **int** type.  To assign **c**, you can use the following code,

```go
c = &a
```

You can also create a pointer **c** by using the short declaration form,

```go
c := &a
```

So, the following code,

```go
fmt.Printf("%p", c)
```

will also output the following memory address on my machine.

```
0xc0000aa058
```

## Why use Pointers?

Pointers are often used when you don't want to make copies of data.  For example, in the following code, when passing the variable **value** into the function **addTwo**, a copy is made and assigned to the variable **n**.

```go
func main() {

	value := 3

	fmt.Println(addTwo(value))
	fmt.Printf("value: %v\n", value)
}

func addTwo(n int) int {
	return n + 2
}
```

In this case **value** is not affected by the function call, as shown by this output.

```
5
value: 3
```

This is an example of the concept of **immutability**.  We have created a function that does not change the value of the original variable.  Another way of creating this function would be as follows,

```go
func main() {

	value := 3

	addTwoPointer(&value)
	fmt.Printf("value: %v\n", value)
}

func addTwoPointer(n *int) {
	*n += 2
}
```
