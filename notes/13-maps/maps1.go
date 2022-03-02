package main

import "fmt"

func main() {
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

	for k := range capital {
		fmt.Printf("The capital of %s is %s.\n", k, capital[k])
	}
}
