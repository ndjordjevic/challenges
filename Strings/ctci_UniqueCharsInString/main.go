package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniqueCharsInString4("nerad"))
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

func uniqueCharsInString2(input string) bool {
	var checker int

	for _, char := range input {
		val := char - 'a'

		if checker&(1<<val) > 0 {
			return false
		}

		checker |= 1 << val
	}

	return true
}

func uniqueCharsInString3(input string) bool {
	arr := [26]int{}
	for _, c := range input {
		if arr[c-'a'] == 0 {
			arr[c-'a']++
		} else {
			return false
		}
	}

	return true
}

func uniqueCharsInString4(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}

	return true
}
