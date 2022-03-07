package main

import "fmt"

func main() {
	dna := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"

	nucleoMap := numberNucleotides(dna)

	for k := range nucleoMap {
		fmt.Printf("%c : %d\n", k, nucleoMap[k])
	}
}

func numberNucleotides(strand string) map[rune]int {

	nucleo := make(map[rune]int)

	// ADD CODE HERE

	return nucleo
}
