package main

import "fmt"

type Group []Player

func (g Group) String() string {
	group := ""
	for _, s := range g {
		group += fmt.Sprintf("%v\n", s)
	}
	return group[:len(group)-1]
}
