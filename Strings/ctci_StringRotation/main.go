package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isRotation("waterbottle", "erbottlewat"))
}

func isRotation(s1, s2 string) bool {
	l := len(s1)

	if len(s2) == l && l > 0 {
		s1s1 := s1 + s1
		return strings.Index(s1s1, s2) != -1
	}

	return false
}
