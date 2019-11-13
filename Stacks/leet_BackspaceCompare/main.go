package main

import (
	"fmt"
)

func main() {
	fmt.Println(backspaceCompare2("ab#c", "ad#c"))
}

func backspaceCompare(S string, T string) bool {
	s1 := make([]rune, 0)
	s2 := make([]rune, 0)

	for _, v := range S {
		if v != '#' {
			s1 = append(s1, v)
		} else {
			if len(s1) > 0 {
				s1 = s1[:len(s1)-1]
			}
		}
	}

	for _, v := range T {
		if v != '#' {
			s2 = append(s2, v)
		} else {
			if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			}
		}
	}

	return string(s1) == string(s2)
}

func backspaceCompare2(S string, T string) bool {
	s := helper(S)
	t := helper(T)
	return s == t
}

func helper(S string) string {
	tmp := []byte{}
	for i := 0; i < len(S); i++ {
		if S[i] != '#' {
			tmp = append(tmp, S[i])
			continue
		}
		if len(tmp) > 0 {
			tmp = tmp[:len(tmp)-1]
		}
	}
	return string(tmp)
}
