package main

import "fmt"

func main() {

	marks := []int{94, 84, 73, 88, 92, 77, 82}

	fmt.Printf("Average: %0.2f", mean(&marks))
}

func mean(grades *[]int) float64 {
	sum := 0
	for _, v := range *grades {
		sum += v
	}
	return float64(sum) / float64(len(*grades))
}
