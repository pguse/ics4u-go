package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Chess Project!\n")
	chessGroup := loadPlayers("students.txt")

	//fmt.Println(chessGroup)

	games := createGames(chessGroup)
	displayGames(games)

	savePlayers("output.txt", chessGroup)

}

func loadPlayers(filename string) Group {
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

func savePlayers(filename string, g Group) {
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

func displayGames(games []Game) {
	for i, g := range games {
		fmt.Printf("Game %d: %v\n", i+1, g)
	}
}
