package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseString("nenad"))
}

func reverseString(word string) string {
	//var a strings.Builder
	//
	//for i := len(word) - 1; i >= 0; i-- {
	//	a.WriteString(string(word[i]))
	//}
	//
	//return a.String()

	var result = ""
	for _, c := range word {
		result = string(c) + result
	}
	return result
}
