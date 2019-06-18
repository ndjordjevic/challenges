package main

import "fmt"

func main() {
	fmt.Println(Maps([]int{4, 1, 2}))
}

func Maps(x []int) []int {
	result := make([]int, len(x))
	for i, v := range x {
		result[i] = v + v
	}
	return result
}
