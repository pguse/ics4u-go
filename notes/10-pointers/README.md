# Pointers

## Variables

Variables in Go have four properties:  a name, a type, a value, and a memory address _(where the value is stored in RAM)_.  You can access the **value** as follows,

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

If you run this on your machine or I run it another time on my machine, the memory address will be different. Notice the use of the **&** character in front of the variable **a**.  The symbol **&a** is a pointer to the variable **a** of type **\*int**.  Notice that the code,

```go
fmt.Printf("%T", &a)
```

outputs the type

```
*int
```

## Pointer Variables

A pointer stores the **memory address** in RAM of a variable.  To declare a pointer, you can do the following,

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

	fmt.Println(value)
	addTwoPointer(&value)
	fmt.Printf("value: %v\n", value)
}

func addTwoPointer(n *int) {
	*n += 2
}
```

In this case, notice that there is no **return** value for the function.  It is unnecessary because by passing a pointer to the original variable into the function, the value at that address is being modified directly.  So, the variable **value** is being directly affected by the actions of the function **addTwoPointer**.  You can see this by looking at the output.

```
3
value: 5
```

## Pointers - Array Storage

By using pointers, you can see that when the elements of an **array** are stored in memory _(RAM)_, they are stored in blocks of memory that are **side by side**.  The following code demonstrates this.

```go
n := [5]int{2, 4, 8, 16, 32}

// array elements are stored in adjacent memory addresses
fmt.Printf("%p, %p, %p, %p, %p\n", &n[0], &n[1], &n[2], &n[3], &n[4])

// note: each int takes up 64 bits or 8 bytes of memory (RAM)
fmt.Println(bits.UintSize)
```

In Go the **int** data type uses **64 bits or 8Mb of memory** to store each integer value.  You can see this by looking at the output of this code.

```
0xc00000a3c0, 0xc00000a3c8, 0xc00000a3d0, 0xc00000a3d8, 0xc00000a3e0
64
```

**Note:**  Of course, the memory addresses will be different each time the program is run.

## Memory Addresses


| 0xc000000000 | 0xc000000001 | 0xc000000002 | 0xc000000003 | 0xc000000004 | 0xc000000005 | 
| :----------: | :----------: | :----------: | :----------: | :----------: | :----------: |
| 00000001     | 00000010     | 00000011     | 00000100     | 00000110     | 00000111     |

Memory addresses are represented in base-16 hexadecimal format, starting with 0x.  Each address stores 1 byte of information or 8 bits.