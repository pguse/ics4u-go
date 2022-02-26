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


Here a new map has been created using a **map literal**, where the map is populated with a set of **key/value** pairs.

A new map can also be created using the built-in function **make**,

```go
ages := make(map[string]int) // mapping from strings to ints
```