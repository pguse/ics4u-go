# Structs

At this point we have used functions to organize code and built-in types to organize data.  Now, we are going to build our own types to organize both code and data.

## Composite Types

We are going to build a **composite type** that describes a point in 2-dimensional space. In Go, a programmer described composite type is called a **struct**.  A struct is made up of **fields** that may contain different types.  A **struct** definition for a **point** would look like this

```go
type point struct {
	x float64
	y float64
}
```

Once a **struct** type is declared, we can define variables of that type

```go
var p1 point
```

where each **field** is set to its **default** value.  In this case, both fields will be set to a value of **0.0**.  You can then set the **fields** by using **dotted notation**.  For example,

```go
p1.x = 4
p1.y = 5
```

You can also initialize a variable of a **struct** type using a **struct literal** as follows,

```go
p2 := point{1, 5}
```

When using a **struct literal** you must always pass in the field values in the **same order** in which they are declared in the **struct**.

You can also use  the **name: value** syntax for initializing a **struct**

```go
p3 := point{x: 7, y: 9}
```

_(the order of fields is irrelevant when using this syntax)_. Note:  This allows you to initialize only a **subset of fields**. All the uninitialized fields are set to their **default** zero values.

To access individual fields of a **struct** you have to use **dotted notation**.  For example,

```go
fmt.Printf("Point #1: (%0.1f, %0.1f)\n", p1.x, p1.y)
```

produces the output,

```
Point #1: (4.0, 5.0)
```

If you want to use a function with a **struct** type, it can be used like any other type.  For example, the function **midpoint**

```go
func midpoint(p1, p2 point) point {
	return point{(p1.x + p2.x) / 2, (p1.y + p2.y) / 2}
}
```

takes two parameters of type **Point** and returns a value of type **Point**.  The following code,

```go
fmt.Printf("Midpoint: %v\n", midpoint(p1, p2))
```

produces the output,

```
Midpoint: {5.5 7}
```

## Types and Methods

Using a **struct** you can create a **user-defined** type.  Like other modern languages **Go** allows you to attach **methods** to these types.  For example, to declare a **Fraction** type, we could used the following code,

```go
type Fraction struct {
	Num int
	Den int
}
```

Method declarations look just like function declarations, with one addition:  the **receiver** specification. The **receiver** appears between the keyword **func** and the name of the method.  Just like all other variable declarations, the **receiver** name appears before the type.  For example,

```go
func (f Fraction) String() string {
	return fmt.Sprintf("%d / %d", f.Num, f.Den)
}
```

This method can be used as follows,

```go
f1 := Fraction{2, 3}
fmt.Println(f1.String())
```

by using the **dotted notation**.  This gives the output,

```
2 / 3
```

Note:  The function **Sprintf** returns a **string** using a **format specifier**.  See the [documentation](https://pkg.go.dev/fmt?msclkid=870687ebaefe11ecaae0dbc15511b792#Sprintf).

Just like functions, **methods** can also take parameters of the user-defined type.  For example,

```go
func (f Fraction) Mult(f1 Fraction) Fraction {
	return Fraction{f.Num * f1.Num, f.Den * f1.Den}
}
```

This method can be used to multiply two fractions, as follows

```go
f1 := Fraction{2, 3}
f2 := Fraction{4, 5}
f3 := f1.Mult(f2)
fmt.Println(f3.String())
```

Notice how the **dotted notation** is used along with the argument provided to the **Mult** function.  The above code produces the following output,

```
8 / 15
```