package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(twoCharacters("asdcbsdcagfsdbgdfanfghbsfdab"))
}

func twoCharacters(s string) int32 {
	t := 0
	unique := make(map[rune]int)

	for i := range s {
		unique[rune(s[i])] = unique[rune(s[i])] + 1
	}
	fmt.Println("unique:", unique)

	keys := make([]rune, 0, len(unique))
	for k := range unique {
		keys = append(keys, k)
	}

	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			fmt.Println(string(keys[i]), string(keys[j]))
			fmt.Println(removeBut(keys[i], keys[j], s))
			rem := removeBut(keys[i], keys[j], s)

			l := len(rem)

			if isAlt(rem) {
				if l > t {
					t = l
					fmt.Println("max alt", rem)
				}
			}
		}
	}

	return int32(t)
}

func isAlt(s string) bool {
	if len(s) < 2 {
		return false
	}

	f := s[0]
	sec := s[1]

	if f == sec {
		return false
	}

	for i := 2; i < len(s); i += 2 {
		if s[i] != f || i+1 < len(s) && s[i+1] != sec {
			return false
		}
	}

	return true
}

func removeBut(r rune, r2 rune, s string) string {
	var ret strings.Builder

	for i := range s {
		if rune(s[i]) == r || rune(s[i]) == r2 {
			ret.WriteString(string(s[i]))
		}
	}

	return ret.String()
}
