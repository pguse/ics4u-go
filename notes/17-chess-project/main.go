package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "students.txt"

func main() {

	chessGroup := loadPlayers()

	// Create games between players with adjacent rankings
	games := createGames(chessGroup)

	menu(games, chessGroup)

}

func menu(games []Game, chessGroup Group) {
	var option int
	fmt.Println("Welcome to the Chess Project!")
	fmt.Println("\n1 - Enter Game Result")
	fmt.Println("2 - Quit")

	fmt.Print("\nEnter an option: ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		updateGames(games, chessGroup)
		fmt.Println()
		menu(games, chessGroup)
	case 2:
		// Sort the players in chessGroup by rank
		sort.Sort(chessGroup)
		savePlayers(chessGroup)
	}
}

func loadPlayers() Group {
	g := Group{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return g
	}

	lines := strings.Split(string(data), "\r\n")

	for _, student := range lines {
		s := strings.Split(student, ",")
		rank, _ := strconv.Atoi(s[3])
		g = append(g, Player{s[0], s[1], s[2], rank})
	}

	return g
}

func savePlayers(g Group) {
	f, err := os.Create(filename) // creates a file at current directory
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	f.WriteString(g.String())
}

func createGames(g Group) []Game {
	games := []Game{}
	for i := 0; i < len(g)-1; i = i + 2 {
		games = append(games, Game{g[i], g[i+1]})
	}
	return games
}

func updateGames(games []Game, chessGroup Group) {
	var gameNumber int
	var winner string

	displayGames(games)

	fmt.Print("\nEnter the game to update: ")
	fmt.Scanln(&gameNumber)

	fmt.Printf("Winner of game %d: ", gameNumber)
	fmt.Scanln(&winner)
	// update the rankings
	//  - assume p1 in each game is the higher ranked
	//  - only update rankings if the lower ranked player wins
	index := (gameNumber - 1) * 2 // Game #1 -> index=0; Game #2 -> index-2, etc.
	if games[gameNumber-1].p2.First == winner {
		// determine the index of p1 in our chessGroup
		// based on the game number
		swap := chessGroup[index].Rank
		chessGroup[index].Rank += chessGroup[index+1].Rank - 400
		chessGroup[index+1].Rank += swap + 400
	} else {
		swap := chessGroup[index].Rank
		chessGroup[index].Rank += chessGroup[index+1].Rank + 400
		chessGroup[index+1].Rank += swap - 400
	}
}

func displayGames(games []Game) {
	fmt.Println()
	for i, g := range games {
		fmt.Printf("Game %d: %v\n", i+1, g)
	}
}
