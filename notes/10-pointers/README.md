# Pointers

## Variables

Variables in Go have four properties:  a name, a type, a value, and a memory address _(where the value is stored in RAM)_.  When you declare a variable - telling the compiler what type it is and what it is called - the declaration enables the compiler to set aside a block of memory to store the variable.  For every variable you declare, the compiler remembers which block of memory is used to store the variable.  You can access the **value** of the variable as follows,

```go
var a int = 52
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

To access the start of the memory address that holds the integer value **a**, the following code

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

I **do not recommend** creating functions like this, in most cases.  Instead of **dereferencing** a pointer _(e.g. *n)_ , to change the value of a variable outside the function, just use a **return** statement to return a copy, as in the first example.

## Pointers - Array Storage

By using pointers, you can see that when the elements of an **array** are stored in memory _(RAM)_, they are stored in blocks of memory that are **side by side**.  The following code demonstrates this.

```go
n := [5]int{2, 4, 8, 16, 32}

// array elements are stored in adjacent memory addresses
fmt.Printf("%p, %p, %p, %p, %p\n", &n[0], &n[1], &n[2], &n[3], &n[4])

// note: each int takes up 64 bits or 8 bytes of memory (RAM)
fmt.Println(bits.UintSize)
```

In Go the **int** data type uses **64 bits or 8 bytes of memory** to store each integer value.  You can see this by looking at the output of this code.  Each **int** element of the array uses 8 memory addresses to store it _(8 x 8 bits = 64 bits)_.

```
0xc00000a3c0, 0xc00000a3c8, 0xc00000a3d0, 0xc00000a3d8, 0xc00000a3e0
64
```

**Note:**  Of course, the memory addresses will be different each time the program is run.

## Memory Addresses


| 0x00000000 | 0x00000001 | 0x00000002 | 0x00000003 | 0x00000004 | 0x00000005 | 0x00000006 | 0x00000007 |
| :----------: | :----------: | :----------: | :----------: | :----------: | :----------: | :----------: | :----------: |
| 00000001     | 00000010     | 00000011     | 00000100     | 00000110     | 00000111     | 00001000     | 00001001     |

Memory addresses are represented in base-16 hexadecimal format, starting with 0x.  Each address stores **1 byte** of information or **8 bits**.  In Go, on a computer with a 64-bit processor, **integers** are stored using **64 bits** of memory.  In other words, 8 memory addresses _(8 bits each)_ are used to store each **int** value.

## Binary Numbers

We are accustomed to working with base-10 numbers.  These values are constructed, with **digits 0-9**, using place values that are **powers of 10**.  For example,

| 1000 | 100 | 10  | 1   |
| :--: | :-: | :-: | :-: |
|  4   |  3  |  1  |  5  |

shows that the value **4315** is really **4**x1000 + **3**x100 + **1**x10 + **5**x1.

Binary values are constructed, with **digits 0 and 1**, using place values that are **powers of 2**.  For example, the binary number **1011**

| 8    | 4   | 2   | 1   |
| :--: | :-: | :-: | :-: |
|  1   |  0  |  1  |  1  |

represents the value **1**x8 + **0**x4 + **1**x2 + **1**x1 = **11** in base-10.

## Hexadecimal Numbers

Hexadecimal values are constructed, with **digits 0-9** and **characters a-f**, using place values that are **powers of 16**.  Note:  The characters **a-f** represent the base-10 values **10-15**.  For example, the hexadecimal number _(starting with 0x)_ **0x01a9**

| 4096 | 256 | 16   | 1  |
| :--: | :-: | :-: | :-: |
|  0   |  1  |  a  |  9  |

represents the value **0**x4096 + **1**x256 + **10**x16 + **9**x1 = **425** in base-10.

In **Go**, you can use the **Printf** function to output values in different formats.  For example, you can use the **format verbs**:  **%b** for binary, and **%x** for hexadecimal values.

```go
fmt.Printf("%b\n", 11)
fmt.Printf("%x", 425)
```

produces the output

```
1011
1a9
```

In **Go** the leading characters **0b** represent a binary value while **0x** represent a hexadecimal value, so

```go
fmt.Printf("%d\n", 0b1011)
fmt.Printf("%d", 0x1a9)
```

produces the output

```
11
425
```