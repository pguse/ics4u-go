package main

import "fmt"

type Player struct {
	Last  string
	First string
	House string
	Rank  int
}

func (s Player) String() string {
	return fmt.Sprintf("%s,%s,%s,%d", s.Last, s.First, s.House, s.Rank)
}
