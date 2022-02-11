package main

import "fmt"

func main() {
	board := [9]string{"X", "O", "X", "O", "X", "-", "-", "X", "O"}

	display(board)
}

func display(b [9]string) {
	for i := 0; i < len(b); i += 3 {
		fmt.Printf("%s  %s  %s\n", b[i], b[i+1], b[i+2])
	}
}

func isWin(b [9]string) bool {
	return false
}

func isTie(b [9]string) bool {
	return false
}
