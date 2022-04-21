package main

import "fmt"

type Game struct {
	p1 Player
	p2 Player
}

func (g Game) String() string {
	return fmt.Sprintf("a: %-10s   b: %-10s", g.p1.First, g.p2.First)
}
