package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Printf("%s", data)
	//fmt.Println(data)
	fmt.Println(string(data))
}
