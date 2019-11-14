package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(J string, S string) int {
	var r int

	for _, v := range J {
		for _, v1 := range S {
			if v == v1 {
				r++
			}
		}
	}

	return r
}
