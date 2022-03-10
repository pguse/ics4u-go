# Structs

At this point we have used functions to organize code and built-in types to organize data.  Now, we are going to build our own types to organize both code and data.

## Composite Types

We are going to build a **composite type** that describes a point in 2-dimensional space. In Go, a programmer described composite type is called a **struct**.  A struct is made up of **fields** that may contain different types.  A **struct** definition for a **point** would look like this"

```go
type Point struct {
	x float64
	y float64
}
```

You can initialize a variable of a struct type using a **struct literal** as follows,

```go
p1 := Point{4, 5}
```

When using a **struct literal** you must always pass in the field values in the **same order** in which they are declared in the struct.

You can also use  the **name: value** syntax for initializing a **struct**

```go
p2 := Point{x: 7, y: 9}
```

_(the order of fields is irrelevant when using this syntax)_. Note:  This allows you to initialize only a subset of fields. All the uninitialized fields are set to their corresponding zero value.

To access individual fields of a struct you have to use dot (.) operator.  For example,

```go
	fmt.Printf("Point #1: (%0.1f, %0.1f)\n", p1.x, p1.y)
```

produces the output,

```
Point #1: (4.0, 5.0)
```

If you want to use a function with a **struct** type, it can be used like any other type.  For example, the function **midpoint**

```go
func midpoint(p1, p2 Point) Point {
	return Point{(p1.x + p2.x) / 2, (p1.y + p2.y) / 2}
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