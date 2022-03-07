# Using Maps in Go

In Visual Studio Code, create a folder called **Maps** and then open it. Now open a new Terminal window. 

## Exercises

## 09-0:  Number of Vowels

Modify the starter code called **numberVowelsMaps.go** so that it counts the **number of vowels** found in a text file called **emma.txt** .  You will need to save the text file **emma.txt** in the same folder as **numberVowels.go**.  This is similar to a previous exercise, except now we are using a **map** to store the data for each vowel.

## 09-1:  Number of Nucleotides

The genetic language of every living thing on the planet is DNA. DNA is a large molecule that is built from an extremely long sequence of individual elements called nucleotides. 4 types exist in DNA and these differ only slightly and can be represented as the following symbols: **A** for adenine, **C** for cytosine, **G** for guanine, and **T** for thymine."

Given a single stranded DNA string, compute how many times each nucleotide occurs in the string.

In the starter code, **dna.go**, modify the function called **numberNucleotides(strand string) map[rune]int** that returns a map with keys that are the nucleotides:  **'A'**, **'C'**, **'G'**, and **'T'** and values that represent **how many times each nucleotide occurs** in a string called **strand**.