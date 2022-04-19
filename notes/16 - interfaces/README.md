# Interfaces

## Key Ideas

* How can we use interfaces found within a package?
* Why would we want/need to use an interface?

An **interface** in our daily lives is what we are presented with, on the surface, that allows us to interact with a machine or piece of technology.  For example, your phone's interface consists of a touchscreen and icons.  A modern TV has an interface that consists of a screen of icons and a remote.  A radio in a car may have an interface that consists of physical dials.

When we use an interface, we don't need to know how the object in question works under the surface.  We only need to know **what the interface consists of** and **how to operate its parts**.

In **Go**, an **interface within a package** just consists of **a set of method names**.  If you create a new **type**, and provide it with methods of the same name as an interface, then a particular package knows how to deal with that new type.  In the note below, you will read how the **Stringer** interface in the **fmt** package allows you to teach **Go** how to output your new type in the console with any method within the **fmt** package - e.g. **Print**, **Printf**, **Println**, **Sprintf**, etc.

We will deal with this topic in a limited way this year.  An **interface** is an **abstract type** that we are going to encounter in a couple examples.

## Example #1 - The Stringer Interface

We have been using the **fmt** package when we need to produce output to the console using the  **Print**, **Println**, and **Printf** methods.  Within the **fmt** package there is an **abstract type** called the **Stringer** interface.  Here is what it looks like,

```go
type Stringer interface {
    String() string
}
```

How does this affect us?  Well, for any **type** that we create, if we include a **String** method with a **receiver** of that type, then any method from the **fmt** package will be able to output our type automatically.  For example, we have just recently created a **Card** type that looks like this,

```go
type Card struct {
	Rank string // Ace, Queen, 9, etc.
	Suit rune   // Diamond, Spade, etc
}

func (c Card) String() string {
	return fmt.Sprintf("%s%c", c.Rank, c.Suit)
}
```

We have learned that to use a **method** that we create, the **dotted notation** is used as demonstrated in the following example,

```go
c := Card{"A", '♣'}
fmt.Printf("Here is a card: %v\n", c.String())
```

This produces the output,

```
Here is a card: A♣
```

But, the **fmt** package actually already knows how to output an **instance** of our **Card** type, since we have included a **String** method.  So, this code

```go
c := Card{"A", '♣'}
fmt.Printf("Here is a card: %v\n", c)
```

also produces the output,

```
Here is a card: A♣
```

without using the **dotted notation** to run the **String** method on the **Card** instance **c**.

## Example #2: Sort

**Go** has a standard **sort** package for doing sorting.  The sorting functions sort data **in-place**, which means they do not return a copy of the sorted sequence.

### Integers

Here is an example of how to **sort** a **slice of integers**,

```go
import (
	"fmt"
	"sort"
)

func main() {
	values := []int{2, -5, 8, -10, 3, 6, 0, 15, 11}
	sort.Ints(values)
	fmt.Println(values)
}
```

in **ascending order**.  The output is,

```
[-10 -5 0 2 3 6 8 11 15]
```

### Strings

If you want to **sort** a **slice of strings**, the following code provides an example,

```go
values := []string{"ruby", "python", "go", "c", "java", "pyret", "haskell", "racket"}
sort.Strings(values)
fmt.Println(values)
```

with the output,

```
[c go haskell java pyret python racket ruby]
```

The **sort** package contains the **sort.Interface** which looks like this,

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

The **interface** contains three functions:  **Len**, **Less**, and **Swap**.

### Hand

So, lets say we want to be able to sort our **type** called **Hand**.  All we need to do is create these three methods with a **receiver** of type **Hand**.  For example, the following code

```go
var value = map[string]int{"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13}

type Hand []Card

func (h Hand) String() string {
	hand := ""
	for _, c := range h {
		hand += fmt.Sprintf("%v ", c)
	}
	return hand
}

func (h Hand) Len() int {
	return len(h)
}

func (h Hand) Less(i, j int) bool {
	return value[h[i].Rank] < value[h[j].Rank]
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
```

implements the **Len**, **Less**, and **Swap** methods along with a **String** method.  Notice:  to say that a **Card** in a **Hand** is less than another one, the values of their ranks are compared using the **values** map.

The following code provides an example of how to sort an instance of our **Hand** type,

```go
h := Hand{Card{"J", '♣'}, Card{"A", '♠'}, Card{"5", '♢'}, Card{"10", '♡'}, Card{"A", '♡'}}
fmt.Println(h)
sort.Sort(h)
fmt.Println(h)
```

producing the output

```
J♣ A♠ 5♢ 10♡ A♡
A♠ A♡ 5♢ 10♡ J♣
```