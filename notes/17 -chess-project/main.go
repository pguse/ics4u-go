package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Last  string
	First string
	House string
	Rank  int
}

func (s Student) String() string {
	return fmt.Sprintf("%s,%s,%s,%d", s.Last, s.First, s.House, s.Rank)
}

type Class []Student

func (c Class) String() string {
	class := ""
	for _, s := range c {
		class += fmt.Sprintf("%v\n", s)
	}
	return class[:len(class)-1]
}

func main() {

	arts := Class{}

	data, err := ioutil.ReadFile("students.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\r\n")
	fmt.Println("Wecome to the Chess Project!")
	//fmt.Println(lines)

	for _, student := range lines {
		s := strings.Split(student, ",")
		rank, _ := strconv.Atoi(s[3])
		arts = append(arts, Student{s[0], s[1], s[2], rank})
	}

	fmt.Println(arts)

	f, err := os.Create("output.txt") // creates a file at current directory
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	f.WriteString(arts.String())

}
