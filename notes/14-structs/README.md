# Structs

At this point we have used functions to organize code and built-in types to organize data.  Now, we are going to build our own types to organize both code and data.

## Composite Types

We are going to build a **composite type** that describes a point in 2-dimensional space. In Go, a programmer described composite type is called a **struct**.  A struct is like a _blueprint_ that is used for creating **objects** of a particular type.  A **struct** definition for a point would look like this"

```go
type Point struct {
	x float64
	y float64
}
```
