package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(oneWay("pale", "pale"))
}

func oneWay(s1 string, s2 string) bool {
	var ss, bs string

	if math.Abs(float64(len(s1))-float64(len(s2))) > 1 {
		return false
	}

	if len(s1) >= len(s2) {
		bs = s1
		ss = s2
	} else {
		bs = s2
		ss = s1
	}

	m := make(map[rune]int)

	for _, v := range bs {
		m[v] = m[v] + 1
	}

	fmt.Println(m)

	m2 := make(map[rune]int)

	for _, v := range ss {
		m2[v] = m2[v] + 1
	}

	fmt.Println(m2)

	var diff int

	for _, v := range bs {
		if m[v] != m2[v] {
			diff++
		}
	}

	return diff <= 1
}
