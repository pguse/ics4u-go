package main

import "fmt"

type Group []Player

func (g Group) Len() int {
	return len(g)
}

func (g Group) Less(i, j int) bool {
	return g[i].Rank < g[j].Rank
}

func (g Group) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Group) String() string {
	group := ""
	for _, s := range g {
		group += fmt.Sprintf("%v\n", s)
	}
	return group[:len(group)-1]
}
