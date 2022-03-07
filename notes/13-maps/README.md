# Maps

A map contains a collection of indices, which are called **keys**, and a collection of **values**. Each key is associated with a single value. The association of a key and a value is called a **key-value pair**.  Unlike an array, the indices/keys can be of any type.  Hence, whereas an array is an ordered collection of data, a map is a **mapping** from keys to values - there is **no inherent order** to the data.

Below, you can see an example of how to create a map using the province-captial information from Canada.

```go
capital := map[string]string{
		"Ontario":                   "Toronto",
		"Quebec":                    "Quebec City",
		"Nova Scotia":               "Halifax",
		"New Brunswick":             "Fredericton",
		"Manitoba":                  "Winnipeg",
		"British Columbia":          "Victoria",
		"Prince Edward Islan":       "Charlottetown",
		"Saskatchewa":               "Regina",
		"Alberta":                   "Edmonton",
		"Newfoundland and Labrador": "St.John's",
		"Northwest Territories":     "YellowKnife",
		"Yukon":                     "Whitehorse",
		"Nunavut":                   "Iqaluit",
	}
}
```


Here a new map has been created using a **map literal**, where the map is populated with a set of **key/value** pairs.  Unlike an **array** or **slice**, the **key-value** pairs are not saved in any order.  You can demonstrate this by running the following code,

```go
for k := range capital {
		fmt.Printf("The capital of %s is %s.\n", k, capital[k])
	}
```

This could produce the following output

```
The capital of Prince Edward Islan is Charlottetown.
The capital of Alberta is Edmonton.
The capital of Newfoundland and Labrador is St.John's.
The capital of Ontario is Toronto.
The capital of Quebec is Quebec City.
The capital of Nova Scotia is Halifax.
The capital of New Brunswick is Fredericton.
The capital of Manitoba is Winnipeg.
The capital of British Columbia is Victoria.
The capital of Yukon is Whitehorse.
```

However, running the code again will produce the output in a different order, since a **map** is not stored in an ordered format.

A new map can also be created using the built-in function **make**,

```go
ages := make(map[string]int) // mapping from strings to ints
```

Then, **key-value** pairs can be created as follows,

```go
marks["alice"] = 85
marks["charlie"] = 75
marks["bob"] = 90
```

Again, if you try to output the map,

```go
for k := range marks {
	fmt.Printf("%-10s %d\n", k, marks[k])
}
```

it will output in some random order,

```
charlie    75
bob        90
alice      85
```

Note:  Unlike an **array** you can continue to add **key-value** pairs without first specifying a length.  However, you can at any time determine the number of **key-value** pairs in a map using the **len** function,

```go
fmt.Println(len(marks))
```

gives output

```
3
```

## Maps - Ordered Output

In order to output a **map** in order, you can use a **slice** to hold the **keys** and then **sort** the slice.  For example,

```go
namesSlice := make([]string, 0, len(marks))
```

creates a **slice** of strings with a **length** of 0 and a capacity equal to the length of the marks map - **len(marks)**.  Now the **keys** must be appended to the empty slice,

```go
for k := range marks {
	namesSlice = append(namesSlice, k)
}
```

and the slice can be **sorted in place**.

```go
sort.Strings(namesSlice)
```

Looping through the **sorted slice of keys** 

```go
for _, name := range namesSlice {
	fmt.Printf("%-10s %d\n", name, marks[name])
}
```

outputs the map in alphabetical order.

```
alice      85
bob        90
charlie    75
```

# Using a Map as a Counter

You can use a **map** to keep track of the number of occurrences of some type of data.  For example, if we wanted to count the number of emojis in the following string

```go
text := "👽😊🍟😊👽😺👽"
```

we can first create a **map** with **keys** of type **rune** and values of type **int**,

```go
emoji := make(map[rune]int)
```

A **rune** in Go refers to a single character.  The following code, iterates through the string **text** using the **for-range** loop,

```go
for _, em := range text {
	emoji[em] += 1
	fmt.Printf("%c ", em)
}
```

and has an output

```
👽 😊 🍟 😊 👽 😺 👽
```

A map in Go has a default value for any **key**, that depends on the type of **value** in the **map**.  In this case, every possible **key** has a default value of 0 (the default value for **int**).  So, the line

```go
emoji[em] += 1
```

updates the **value** for any **key** regardless of whether the **key** already exists or not.

The following code,

```go
for k := range emoji {
	fmt.Printf("%c : %d\n", k, emoji[k])
}
```

iterates through the **map** using the **for-range** loop and outputs the **key-value** pairs as follows,

```
👽 : 3
😊 : 2
🍟 : 1
😺 : 1
```

The order of the output will vary, because the **key-value** pairs in **maps** are not stored in any particular order.