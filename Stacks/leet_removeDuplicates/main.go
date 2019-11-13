package main

import (
	"fmt"
	"github.com/ndjordjevic/challenges/Stacks/stack"
	"strings"
)

func main() {
	fmt.Println(removeDuplicates3("abbaca"))
}

func removeDuplicates(S string) string {
	r := strings.Builder{}
	stk := stack.New()
	for _, v := range S {
		if stk.Length() == 0 {
			stk.Push(v)
			continue
		}

		if stk.Peek() != v {
			stk.Push(v)
		} else {
			_, _ = stk.Pop()
		}
	}

	for stk.Peek() != nil {
		v, _ := stk.Pop()
		r.WriteRune(v.(rune))
	}

	return reverse(r.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func removeDuplicates2(S string) string {

	var stk []byte

	for i := 0; i < len(S); i++ {
		if len(stk) != 0 && stk[len(stk)-1] == S[i] {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, S[i])
		}
	}

	return string(stk)
}

func removeDuplicates3(S string) string {
	stk := make([]rune, 0)
	for _, c := range S {
		if len(stk) > 0 && stk[len(stk)-1] == c {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, c)
		}
	}

	var result strings.Builder
	for _, c := range stk {
		result.WriteRune(c)
	}
	return result.String()
}
