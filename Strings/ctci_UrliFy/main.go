package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(urliFy("Mr John Smith    ", 13))
	fmt.Println(urliFy2("Mr John Smith"))
}

func urliFy(s string, l int) string {
	a := []rune(s)
	var spaceCount int

	for i := 0; i < l; i++ {
		if a[i] == ' ' {
			spaceCount++
		}
	}

	index := l + spaceCount*2
	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			a[index-1] = '0'
			a[index-2] = '2'
			a[index-3] = '%'
			index -= 3
		} else {
			a[index-1] = a[i]
			index--
		}
	}

	return string(a)
}

func urliFy2(s string) string {
	b := strings.Builder{}

	for _, v := range s {
		if v != ' ' {
			b.WriteRune(v)
		} else {
			b.WriteString("%20")
		}
	}

	return b.String()
}
