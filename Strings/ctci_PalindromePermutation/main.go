package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(palindromePermutation("Tact Coa"))
}

func palindromePermutation(s string) bool {
	table := [26]int{}

	for _, v := range s {
		if unicode.IsLetter(v) {
			table[unicode.ToLower(v)-'a']++
		}
	}

	foundOdd := false

	for _, v := range table {
		if v%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}

	return true
}
