package main

import (
	"fmt"
)

func main() {
	fmt.Println(countCharacters2([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}

func countCharacters(words []string, chars string) int {
	r := 0

	m := map[rune]int{}
	for _, c := range chars {
		m[c] = m[c] + 1
	}

outside:
	for _, v := range words {
		u := map[rune]int{}
		for _, c := range v {
			u[c] = u[c] + 1
		}

		for k, v := range u {
			if m[k] < v {
				continue outside
			}
		}

		r = r + len(v)
	}

	return r
}

func countCharacters2(words []string, chars string) int {
	ccm := charCountMap(chars)
	l := 0
	for _, word := range words {
		if isGood(word, ccm) {
			l += len(word)
		}
	}
	return l
}

func charCountMap(chars string) [26]int {
	ccm := [26]int{}
	for _, c := range chars {
		ccm[c-'a']++
	}
	return ccm
}

func isGood(word string, ccm [26]int) bool {
	wccm := [26]int{}
	for _, c := range word {
		k := c - 'a'
		wccm[k]++
		if wccm[k] > ccm[k] {
			return false
		}
	}
	return true
}
