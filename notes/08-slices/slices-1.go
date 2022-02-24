package main

import "fmt"

func main() {
	numberSlice := []int{87, 65, 92}
	numberArray := [4]int{54, 78, 23, 87}

	numberSlice = append(numberSlice, 99)

	// Calculating the Average
	fmt.Printf("Average: %0.2f\n", average(numberSlice))

	// Working with Slices
	fmt.Println("Doubling a Slice")
	fmt.Println(numberSlice)
	doubleSlice(numberSlice)
	fmt.Println(numberSlice)

	fmt.Println("Doubling an Array")
	fmt.Println(numberArray)
	doubleArray(numberArray)
	fmt.Println(numberArray)

}

func doubleSlice(n []int) {
	for i := 0; i < len(n); i++ {
		n[i] *= 2
	}
}

func doubleArray(n [4]int) {
	for i := 0; i < len(n); i++ {
		n[i] *= 2
	}
}

func average(n []int) float64 {
	s := 0
	for i := 0; i < len(n); i++ {
		s = s + n[i]
	}

	return float64(s) / float64(len(n))
}
