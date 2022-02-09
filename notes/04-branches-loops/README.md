# Branches and Loops

# Making Decisions

All programs have been executed sequentially, line by line, up to this point.  Often programs need to make decisions based on certain conditions determined by the value of a variable.  In Go, decisions are made using the if / else if /else statement.

## Comparison Operators

These statements require that you use one of a number of comparison operators.  For example,

* \<   ***(less than)*** the condition is true if left side is less than right side. 
* \>   ***(greater than)*** the condition is true if left side is greater than right side. 
* \<\=  ***(less than or equal)***
* \>\=  ***(greater than or equal)***
* \=\=  ***(equal to)***
* \!\=  ***(not equal to)***

Here are some examples using and **if/else** statement in Go:

```go
n = 17

if n % 2 == 0 {
    fmt.Println("The number is even.")
} else {
    fmt.Print("The number is odd.")
}
```

If there are more than two options, then the **if/else if/else** could be used,

```go
mark := 72

if mark >= 80 {
    fmt.Println("A")
} else if mark >= 70 {
    fmt.Println("B")
} else if mark >= 60 {
    fmt.Println("C")
} else if mark >= 50 {
    fmt.Println("D")
} else {
    fmt.Println("You are not meeting course expectations.")
}
```

In both examples above, only one of the available options will be acted upon in an **if/else if/else** block of code.  You can have as many **else if's** as you need ***(including none)*** .  You should at least have an **if** and **else** if there are only two options, but even the **else** can be omitted.

**Notice:**  The position of the curly braces is important.  One must be on the **same line** as the if statement, otherwise the Go compiler will throw an error.  Similarly, the position of the curly braces on the lines with **else if** and **else** must be just like as shown above, otherwise an error will be produced.

## Boolean Values & Logical Operators

All **if/else if/else** statements use **Boolean** values, even though that may not be apparent.  For example,

* fmt.Println( 3 > 1) produces the value **true**
* fmt.Println( 5 < 2 ) produces the value **false**.

For example, the code

```go
fmt.Println(3 > 1)
fmt.Println(5 < 2)
```

produces

```
true
false
```

The values **true** and **false** are called **Boolean** values.

## Using Logical Operators

If you need to test more than one condition you may need to use a **logical operator**.  In Go, the logical operators are the symbols **&&** and **||**.  For example,

* fmt.Println( 3 > 1 **||** 5 < 2) produces the value **true** because one of the conditions is true.  Both conditions must be false to produce the value **false**.
* fmt.Println( 1 < 3 **&&** 5 < 2) produces the value **false** because one condition is false.  Both conditions must be true to produce the value **true**.

For example, the code

```go
fmt.Println(3 > 1 || 5 < 2)
fmt.Println(1 < 3 && 5 < 2)
```

produces

```
true
false
```

These logical operators can be used in **if** statements to simplify your Go code.  For example, let's say we want to determine if **at least one** of two integers is odd.  There are two ways we could write the program,

```go
m := 5
n := 8

if m%2 != 0 {
    fmt.Println("Odd")
} else if n%2 != 0 {
    fmt.Println("Odd")
} else {
    fmt.Println("Neither one of the values is odd.")
}
```

or to be more concise we could use the **||** logical operator,

```go
if m%2 != 0 || n%2 != 0 {
    fmt.Println("Odd")
} else {
    fmt.Println("Neither one of the values is odd.")
}
```

If we wanted to determine if **both integers** are odd, we could use the **&&** logical operator,

```go
if m%2 != 0 && n%2 != 0 {
    fmt.Println("Both Odd")
} else {
    fmt.Println("One of the numbers is even.")
}
```

# Loops

## The for-loop

Unlike a number of programming languages, Go has just one type of loop; the **for-loop**.

```go
for i := 0; i < 3; i++ {
    fmt.Println("Hello everyone!")
}
fmt.Println("Loop finished.")
```

produces the output:

```
Hello everyone!
Hello everyone!
Hello everyone!
Loop finished
```

The variable **i** is initialized to the value 0 before the loop starts.  Then, the condition **i < 3** is evaluated.  If it is **true** the loop starts.  At the end of one iteration the statement **i++** is executed.  This increases the value of **i** by 1.  Next, the condition **i < 3** is evaluated and the loop continues with the next iteration if it is still **true**.

Notice how the value of **i** changes

```go
for i:=0; i < 3; i++ {
    fmt.Println(i, "- Hello everyone!")
}
```

produces the output:

```
0. Hello everyone!
1. Hello everyone!
2. Hello everyone!
```

We can also count backwards by doing the following:

```go
for i:=3; i > 0; i-- {
    fmt.Println(i, "- Hello everyone!")
}
```

produces the output:

```
3 - Hello everyone!
2 - Hello everyone!
1 - Hello everyone!
```

**Notice:** The position of the curly braces is important.  There must be one on the same line as **for** otherwise the Go compiler with throw and error.

## The for-loop - Using an Accumulator

An accumulator is a variable that is used in a loop to construct a value using an iterative process.  It could act like a counter using addition, or perform exponentiation through repeated multiplication, or build a string character by character.  Here is an example of an accumulator in action.

### The Sum of all Integers

The act of adding consecutive integers e.g. 0 + 1 + 2 + 3 + 4 + 5 + ... can be thought of as an iterative process.  In each step you are adding ... the only thing that changes is the number being added.  In Go, this can be done using a loop, as follows:

```go
N := 5
sum := 0
for i := 0; i <= N; i++ {
    sum = sum + i
}
fmt.Println(sum)
```

The result \(15\) that is printed is the sum of all integers between 1 and 5.  Here is a table that summarizes how each variable changes values are the loop iterates.

| N | i | sum |
| :---: | :---: | :---: |
| 5 | - | 0 |
| 5 | 1 | **1** |
| 5 | **2** | **3** |
| 5 | 3 | 6 |
| 5 | 4 | 10 |
| 5 | 5 | 15 |

The assignment statement **sum = sum + i** is the key line and most importantly it is read from **right-to-left**.  The current values of **sum** and **i** are added together and the result stored back in the variable/identifier called **sum**.  In the table above one of these calculations is shown using the bolded values, where **i** and **sum** have the values 2 and 1.  They are added to give the new value of **sum**, which is 3.

## Other forms of the for-loop

You can create a for-loop without the three parts shown above.  For example,

```go
n := 15
for n != 1 {
    if n%2 == 0 {
        n = n/2
    } else {
        n = 3*n + 1
    }
    fmt.Println(n)
}
```

produces the output:

```
46
23
70
35
106
53
160
80
40
20
10
5
16
8
4
2
1
```

You can even leave out all of the parts.  For example,

```go
import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		n := rand.Intn(3)

		fmt.Println(n)

		if n == 0 {
			break
		}
	}
}
```

may produce the following output:

```
1
2
2
4
1
3
0
```