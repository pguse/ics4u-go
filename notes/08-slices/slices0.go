package main

import "fmt"

func main() {
	numberSlice := []int{87, 65, 92}
	anotherSlice := []int{15, 76, 44}

	numberSlice = append(numberSlice, 99)
	numberSlice = append(numberSlice, 54, 35)
	fmt.Println(numberSlice)

	result := append(numberSlice, anotherSlice...)
	fmt.Println(result)
}
