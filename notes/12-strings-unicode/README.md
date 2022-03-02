# Strings

Individual letters, digits, and symbols are called **characters**.  When you _string_ togther characters and place them between quotes, it's called a _literal string_.

Literal values wrapped in quotes are inferred to be of type **string**.  For example,


```go
text := "Hello World!"

fmt.Printf("%T", text)
```

produces the output

```
string
```

You can use the **for-loop** we've been using to loop across a **string**.  For example,

```go
for i := 0; i < len(text); i++ {
    fmt.Printf("%v ", text[i])
}
```

outputs the **Unicode code points** of the characters in the **string**.

```
72 101 108 108 111 32 87 111 114 108 100 33
```

These are the **unsigned integers** used to store each character in the string.  Each code point corresponds to **1 byte** of information.  If you want to output the corresponding characters, you need to use the code,

```go
for i := 0; i < len(text); i++ {
	fmt.Printf("%c", text[i])
}
```

to get the output

```
Hello World!
```

## The for-range loop - Multilingual text

The Go language has a more convenient version of the for-loop that uses the **range** keyword.  It has a simpler form and doesn't require the use of an **index** or the **length** of the string. For example,

```go
for i, ch := range text {
    fmt.Println(i, ch)
}
```

produces the output

```
0 72
1 101
2 108
3 108
4 111
5 33
```

Notice that the variable **i** stores the index while **ch** stores the **Unicode code-point**.  If you don't need the index **i**, it should be replaced with the **_** symbol.  For example,

```go
for _, ch := range text {
	fmt.Printf("%c", ch)
}
```

produces the output

```
Hello!
```

Imagine that we add a character, such as an **emoji** to our string.  The following code, using a standard for-loop,

```go
text := "Hello 🐱!"

for i := 0; i < len(text); i++ {
    fmt.Printf("%d %v\n", i, text[i])
}
```

produces the output

```
0 72
1 101
2 108
3 108
4 111
5 32
6 240
7 159
8 144
9 177
10 33
```

If you look carefully at the string "Hello 🐱!", there are only **8 characters**, yet there are 11 code-points in the output.  The **emoji** is actually stored using **4 bytes** of memory and uses 4 code-points to store it instead of one.  Attempting to output each character using the standard for-loop as

```go
for i := 0; i < len(text); i++ {
	fmt.Printf("%d %c\n", i, text[i])
}
```

produces an output as follows and causes the program to hang.

```
0 H
1 e
2 l
3 l
4 o
5
6 ð
7
```

In comparison, the **for-range** loop handles the string without a problem.  For example,

```go
for i, ch := range text {
	fmt.Printf("%d %c\n", i, ch)
}
```

produces the correct output

```
0 H
1 e
2 l
3 l
4 o
5  
6 🐱
10 !
```

See how the **emoji** uses indices 6-9.  The 4 indices correspond to **4 bytes** of information.

Again, without using the index values, the code

```go
for _, ch := range text {
	fmt.Printf("%c", ch)
}
```

produces the output

```
Hello 🐱!
```

It is best practice to use the **for-range** loop when looping across strings.


This note does not just apply to **emojis** within **strings**.  Similar effects occur when you use languages with characters that are not part of the English alphabet.