# Introduction to Go

Go is a compiled programming language, unlike Python which is an interpreted language.  Before you run a Go program, a compiler translates your code into the 1s and 0s that machines use.  It compiles all your code into a single _executable_ for you to run.  During this process, the Go compiler can catch any of your typos or syntax mistakes.

Here is a simple Go program

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

And here is how you would run the program,

![Running a Go Program](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Running a Go Program")

In the code above, the **package** keyword declares the package this code belongs to.  All Go code must belong to a package.  Each package corresponds to a single idea or concept.

The **import** keyword specifies that our code will use the **fmt** package which contains functions for formatted input and output, such as **Println**.  The **func** keyword declares a function called **main**.  The **main** identifier is special because when you run a Go program, it runs the code found in the **main** function.  Without **main**, the Go compiler would report an error.