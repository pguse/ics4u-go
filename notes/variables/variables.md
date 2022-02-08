# Variables

In Go, variables must be declared as a specific type.  There are a number of ways to declare a variable, but here I will focus on just two. For example,

```go
var name string
var age int
```

In this code, a variable called **name** of type **string** is created as well as a variable **age** of type **int** _(an integer)_.  Since no values are assigned, each variable is automatically given the default value for its type.  For strings the default value is **""** and for integers _(int)_ the default value is 0.  You can also assign different values, when the variables are declared, as follows

```go
var name string = "Albert"
var age int = 17
```

**Short declaration** provides an alternative to using the **var** keyword.  For example,

```go
name := "Albert"
age := 17
```

In this code, the Go compiler interprets the type of each variable using the values assigned. So, **name** will be of type **string** and  **age** will be of type **int**.  _Note: Short declaration of variables can only be used inside a function._

The other numerical type we will initially come across is the **floating-point** type.  For example,

```go
mark := 82.5
```
The variable **mark** will be of type **float64**, which just means a decimal _(floating-point)_ value that is stored using 64 bits of memory in RAM.

To see the type of variable, you can use the formatted print function **Printf** from the **fmt** package, and use the format specifier **%T** to output the type of the variable.  Note:  The format verb is substituted for the value _(in this case, its type)_ in the expression provided by the second argument.  If you want to print out the values of each variable using the **Printf** function, you would use **%s** for strings, **%d** for integers, and **%f** for float-point values.  See the example below,

```go
package main

import "fmt"

func main() {
	name := "Albert"
	age := 17
	mark := 82.5

	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of mark: %T\n", mark)

	fmt.Printf("Value of name: %s\n", name)
	fmt.Printf("Value of age: %d\n", age)
	fmt.Printf("Value of mark: %f\n", mark)
}
```
_Note:  Instead of using different **format specifiers** you can also use **%v** which will output the values in a default format_