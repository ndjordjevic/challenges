package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(caesarCipher("abcdefz", 2))
}

func caesarCipher(s string, k int32) string {
	r := []rune(s)

	for i := range r {
		for j := 0; j < int(k); j++ {
			if unicode.IsLetter(r[i]) {
				if r[i] == 'z' {
					r[i] = 'a'
				} else if r[i] == 'Z' {
					r[i] = 'A'
				} else {
					r[i] += 1
				}
			}
		}
	}

	return string(r)
}
