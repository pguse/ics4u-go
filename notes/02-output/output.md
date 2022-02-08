# Output

We will be creating text output using two functions.  **Println** formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a **newline** is appended.  **Printf** formats according to a format specifier and writes to standard output.  Both of these functions are found in the **fmt** package.  Here is an example of the **Println** function works,

```go
name := "Albert College"
fmt.Println("The name of our school is", name)
```

**Output**
```
The name of our school is Albert College
```

Here is an example of how the **Printf** function works,

```go
fmt.Printf("The name of our school is %s.\n", name)
```

**Output**
```
The name of our school is Albert College.
```

The **Printf** function uses a format specifier that describes how to format the output, depending on the variable type and/or desired output.

Here are a few format specifiers:

| Format Specifier | Description                                      |
| :--------------: |:------------------------------------------------:|
| %v               | the value in a default format                    |
| %d               | an integer in base 10                            |
| %f               | decimal point but no exponent, e.g. 123.456      |
| %s               | the uninterpreted bytes of the string or slice   |

Note:  **Printf** does not automatically append a newline to its output.  If you want a newline, you must use **\n** _(called an escape sequence)_ within your string literal.  Notice that if you want to include a **%** in your output it must be typed as **%%**, because of the use of **%** in the the format specifiers.