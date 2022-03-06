# User Input and Working with Files

## Reading Input from the User

The **Scanf** function scans text read from **standard input**, storing successive space-separated values into successive arguments as **determined by the format**. _It returns the number of items successfully scanned._  You can read a single value as follows,

```go
var name string

fmt.Print("Enter your name: ")
fmt.Scanf("%s", &name)
fmt.Printf("Hello %s!", name)
```

producing the following output

```
Enter your name: Paul
Hello Paul!
```

Notice:  The **Print** function is used instead of **Println** because we want the input to occur on the same line as the prompt.  The **Scanf** function requires the **address** of the variable **name** _(a pointer to **name**)_.  However, the input is stored in the variable **name**.  You can also use **Scanf** to read in multiple values separated by spaces, as determined by the format provided.  For example,

```go
var name string
var age int

fmt.Print("Enter your name and age: ")
fmt.Scanf("%s  %d", &name, &age)
fmt.Printf("Hello %s!\n", name)
fmt.Printf("Your are %d years old.", age)
```

produces the following output

```
Enter your name and age: Paul 51
Hello Paul!
Your are 51 years old.
```

If you wish to have user input on **different lines**, you must include the newline character **\n** within the format string of **Scanf**.  For example,

```go
var name string
var age int

fmt.Print("Enter your name: ")
fmt.Scanf("%s\n", &name)
fmt.Print("Enter your age: ")
fmt.Scanf("%d", &age)

fmt.Printf("Hello %s!\n", name)
fmt.Printf("Your are %d years old.", age)
```

produces the output,

```
Enter your name: Paul
Enter your age: 51
Hello Paul!
Your are 51 years old.
```

The **Scanln** function stops scanning at a newline, so you can also use it for input on separate lines.  For example,

```go
var name string
var age int

fmt.Print("Enter your name: ")
fmt.Scanln(&name)
fmt.Print("Enter your age: ")
fmt.Scanln(&age)

fmt.Printf("Hello %s!\n", name)
fmt.Printf("Your are %d years old.", age)
```

produces the output

```
Enter your name: Paul
Enter your age: 51
Hello Paul!
Your are 51 years old.
```

You can also use **Scanln** to scan multiple values, separated by spaces, on a single line.  Remember, scanning will stop once a **newline** is created.  For example,

```go
var first string
var last string
var age int

fmt.Print("Enter your name: ")
fmt.Scanln(&first, &last)
fmt.Print("Enter your age: ")
fmt.Scanln(&age)

fmt.Printf("Hello %s %s!\n", first, last)
fmt.Printf("Your are %d years old.", age)
```

produces the output

```
Enter your name: Paul Guse
Enter your age: 51
Hello Paul Guse!
Your are 51 years old.
```


## Reading an entire File into Memory

Here is an example of code that will read an entire text file **test.txt** into memory. Note:  The text file **test.txt** must be in the same folder as your Go program.

```go
import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Printf("%s", string(data))
}
```

The output is all of the text contained in the file **test.txt**.

```
The quick brown fox
jumps over
the lazy dog
```

Notice that the **ReadFile** function returns two values:  **data** - the contents of the file, and **err** - whether an error is created in trying to read the file.  If the file is found and read, then the value of **err** should be **nil**, and this is handled by the **if-statement** that follows.

Notice also that when **data** is output, it is first converted to a **string**.  The following statement,

```go
fmt.Println(data)
```

produces the output,

```
[84 104 101 32 113 117 105 99 107 32 98 114 111 119 110 32 102 111 120 13 10 106 117 109 112 115 32 111 118 101 114 13 10 116 104 101 32 108 97 122 121 32 100 111 103]
```

because the contents **data** of the file are initially stored as a slice of **bytes** _(uint8 values)_, which represent the **Unicode code-points** of each character.