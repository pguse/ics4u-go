# Working with Strings and Numbers

## Problem - Armstrong Numbers

An Armstrong number is a number that is the sum of its own digits each raised to the power of the number of digits.

For example:

- 9 is an Armstrong number, because 9 = 9^1 = 9
- 10 is not an Armstrong number, because 10 != 1^2 + 0^2 = 1
- 153 is an Armstrong number, because: 153 = 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
- 154 is not an Armstrong number, because: 154 != 1^3 + 5^3 + 4^3 = 1 + 125 + 64 = 190

Write some code to determine whether a number is an Armstrong number.

### Starting the Problem

In order to access the digits of a number in this problem, it is more convenient to store the number as a string.  For example,

```go
var number string = "153"

for i, digit := range number {
	fmt.Println(i, digit)
}
```

The **for-range** loop allows you to loop through any sequence of elements _(string, array, splice, etc)_ without knowing/using the length of the sequence.  It stores two values:  **1.** The index **i** of an element; **2.** The value **digit** of the element.  Surprisingly _(based on your experience with Python)_, the output of the code above is,

```
0 49
1 53
2 51
```

The values of the elements are the numerical values that each **character** is stored as using the **unicode/ascii** system.  We can output the associated characters using the **Printf** function as follows,

```go
fmt.Printf("%d %c\n", i, digit)
```

giving us the output,

```
0 1
1 5
2 3
```

What we are encountering in this problem is a data type in **Go** called a **rune**.  The variable **digit** is a **rune**. It can either be represented as its **unicode** value or its associated **character**.  However, in order to solve this problem, we need to use the numerical value of the digits.  We do this be using a package called **strconv** which handles **string** [conversions](https://pkg.go.dev/strconv).

## The strconv Package

Here is the code that allows us to access the numerical value of each digit,

```go
for i, digit := range number {
	n, err := strconv.Atoi(string(digit))
	if err != nil {
		fmt.Println(err)
		break
	}
	fmt.Println(i, n)
}
```

The **strconv.Atoi** function attempts to convert a **string** to an **int**.  Notice:  **1.**  We must first convert the **rune** called **digit** to a **string**; **2.** The **strconv.Atoi** function returns two values, stored as **n** and **err**.  The variable **n** is the integer value, and **err** stores the error if one occurs in the attempted conversion.  For example, if you attempted to convert ```"apple"``` to an **int** and error would be produced.  We can simplify our code above by replacing variables we don't need with the **underscore** character **_**.  This is a simpler version of the same code,

```go
for _, digit := range number {
	n, _ := strconv.Atoi(string(digit))
	fmt.Println(n)
}
```

and produces the output,

```
1
5
3
```

Now if we wanted to add the digits, we could do so like this,

```go
var number string = "153"
var sum int = 0

for _, digit := range number {
	n, _ := strconv.Atoi(string(digit))
	sum += n
}

fmt.Println(sum)
```

To solve the **Armstrong Number** problem, we want to add the digits, each raised to an exponent equal to the number of digits _(the length of the **string** called **number**)_.

The **for-range** loop needs to change to the following,

```go
for _, digit := range number {
	n, _ := strconv.Atoi(string(digit))
	sum += int(math.Pow(float64(n), float64(len(number))))
}
```

Note:  The **math.Pow** function requires two arguments that are each of type **float64**, and returns a **float64** value.  As a result, we need to use a number of **type conversions**.

If the sum is equal to the value of the original number, then the original number is called an **Armstrong Number**.  We need to make a simple decision, so an **if-statement** is required,

```go
if strconv.Itoa(sum) == number {
	fmt.Println("Armstrong Number!")
} else {
	fmt.Println("Not an Armstrong Number")
}
```

Notice the use of the **strconv.Itoa** function.  It converts an **int** to a **string** in order to make the comparison with the original **string** called **number**.


```go
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var number string = "153"
	sum := 0

	for _, digit := range number {
		n, _ := strconv.Atoi(string(digit)) // converts string to int
		sum += int(math.Pow(float64(n), float64(len(number))))
	}

	// Is this an Armstrong number?
	if strconv.Itoa(sum) == number {
		fmt.Println("Armstrong Number!")
	} else {
		fmt.Println("Not an Armstrong Number")
	}
	fmt.Println(sum)

}

```