# Using Structs in Go

In Visual Studio Code, create a folder called **Structs** and then open it. Now open a new Terminal window.

## Exercises

Modify the starter code called **pointFunctions.go**, in completing the following exercises.

## 10-0:  Distance

Complete the following function, that returns the **distance between two points**.

```go
func distance(p1, p2 Point) float64 {
	return 0.0
}
```

## 10-1:  Slope

Complete the following function, that returns the **slope of a line passing through two points**.

```go
func slope(p1, p2 Point) float64 {
	return 0.0
}
```

Modify the starter code called **fractionMethods.go**, in completing the following exercises.

## 10-2:  Divide

Complete the following method, that returns the **quotient** of **f / f1**, where **f** and **f1** are of type **Fraction**.

```go
func (f Fraction) Div(f1 Fraction) Fraction {
	return Fraction{}
}
```

## 10-3:  Add

Complete the following method, that returns the **sum** of **f + f1**, where **f** and **f1** are of type **Fraction**.

```go
func (f Fraction) Add(f1 Fraction) Fraction {
	return Fraction{}
}
```

## 10-4:  Subtract

Complete the following method, that returns the **difference** of **f - f1**, where **f** and **f1** are of type **Fraction**.

```go
func (f Fraction) Subt(f1 Fraction) Fraction {
	return Fraction{}
}
```

Modify the starter code called **cardGame.go**, in completing the following exercises.

## 10-5: Flush

Create a **boolean** method called **Flush** that determines if a hand of cards is a flush _(all the same suit)_.

```go
func (h Hand) Flush() bool {
	// All cards have the same suit
	return false
}
```

## 10-6: Straight

Create a **boolean** method called **Straight** that determines if a hand of cards is a straight _(sequence of ranks)_.  _**Note:**  Assume the hand of cards is already sorted_.

```go
func (h Hand) Straight() bool {
	// Cards are in sequential order - assume a sorted hand
	return false
}
```

Note:  This function requires that we compare the values of cards.  Since the **rank** is stored as a **string**, this comparison is not obvious.  The **global variable** value, defined at the top of **cardGame.go**,

```go
var value = map[string]int{
	"A":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
}
```

maps the **rank** of a card to its **value**.  This **map** can be used in the **Straight** method.


## 10-7: Straight Flush

Create a **boolean** method called **StraightFlush** that determines if a hand of cards is a straight flush _(sequence of ranks, all the same suit)_.  _**Note:**  Assume the hand of cards is already sorted_."

```go
func (h Hand) StraightFlush() bool {
	// Cards are in sequential order with same suit - assume a sorted hand
	return false
}
```

## 10-8: Pair

Create a **boolean** method called **Pair** that determines if a hand of cards has a single pair

```go
func (h Hand) Pair() bool {
	// Two cards with the same rank
	return false
}
```