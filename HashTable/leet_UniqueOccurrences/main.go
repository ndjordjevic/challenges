package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2}))
}

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	m1 := map[int]bool{}

	for _, v := range arr {
		m[v] = m[v] + 1
	}

	for _, value := range m {
		m1[value] = true
	}

	return len(m) == len(m1)
}
