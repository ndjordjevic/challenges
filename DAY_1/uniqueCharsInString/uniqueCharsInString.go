package main

import "fmt"

func main() {
	fmt.Println(uniqueCharsInString("neadNn"))
}

//func uniqueCharsInString(input string) bool {
//	a := make(map[int32]bool)
//
//	for _, v := range input {
//		a[v] = true
//	}
//
//	return len(input) == len(a)
//}

func uniqueCharsInString(input string) bool {
	var a [128]bool

	for _, v := range input {
		if a[v] {
			return false
		}
		a[v] = true
	}

	return true
}
