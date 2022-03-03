# Tic-Tac-Toe Project

This project demonstrates how you might use a single-dimensional array in Go to store information in a grid-based game like Tic-Tac-Toe . In this case a 9 element array is used to store the 9 possible positions of the tic-tac-toe game. The index values of the array  match the grid positions given in the table below.

|   Tic  |  Tac   |  Toe   |
|:---:|:---:|:---:|
|  0  |  1  |  2  |
|  3  |  4  |  5  |
|  6  |  7  |  8  |


## Board Setup

The **board** array is filled initially with the values ```X```, ```O```, or ```-```.  In this game, this value will represent an empty square and be displayed as ```-```.  See the starter code below.

```go
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
```

## Tasks

### 05-0:  Is there a Win?

Complete the **isWin** function. It should return **true** if there is a win on the board.  Otherwise, it should return **false**. Remember, there are 8 possible ways to win in the tic-tac-toe game:  3 of the same kind _('X' or 'O')_ in any row _(3)_, any column _(3)_, or either diagonal _(2)_.

### 05-1:  Is there a Tie?
Complete the **isTie** function. It should return **true** if there is a tie on the board.  Otherwise, it should return **false**.  A **tie** occurs if the board is filled, with 'X's and 'O's but there is **no win**.
