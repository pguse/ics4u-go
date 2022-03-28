# Working with Files

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