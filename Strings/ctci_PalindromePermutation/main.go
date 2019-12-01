package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(palindromePermutation("Tact Coa"))
}

func palindromePermutation(s string) bool {
	m := make(map[rune]int)
	var l int

	for _, v := range s {
		if v != ' ' {
			m[unicode.ToLower(v)]++
			l++
		}
	}

	var oddC int
	for k, v := range m {
		fmt.Printf("%c, %d", k, v)
		fmt.Println()
		if v%2 != 0 {
			oddC++
		}
	}

	fmt.Println(l)
	fmt.Println(oddC)

	if oddC == 1 {
		return true
	}

	return false
}
