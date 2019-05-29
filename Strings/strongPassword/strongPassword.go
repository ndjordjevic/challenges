package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(minimumNumber(11, "#HackerRank"))
}

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	var r int32

	var l, u, d, s bool

	for _, v := range password {
		if unicode.IsLower(v) {
			l = true
		}

		if unicode.IsUpper(v) {
			u = true
		}

		if unicode.IsDigit(v) {
			d = true
		}
	}

	if strings.ContainsAny(password, "!@#$%^&*()-+") {
		s = true
	}

	fmt.Println(l, u, d, s)

	if !l {
		r++
	}

	if !u {
		r++
	}

	if !d {
		r++
	}

	if !s {
		r++
	}

	if n+r < 6 {
		r = 6 - n
	}

	return r
}
