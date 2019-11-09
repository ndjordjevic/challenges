package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit2("RLLLLRRRLR"))
}

func balancedStringSplit(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == s[0] {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}

func balancedStringSplit2(s string) int {
	// RLLLLRRRLR
	res, count := 0, 0
	for _, b := range s {
		if b == 'L' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}
