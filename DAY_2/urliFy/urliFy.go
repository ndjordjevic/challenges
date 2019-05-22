package main

import "fmt"

func main() {
	fmt.Println(urliFy("Mr John Smith    ", 13))
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
