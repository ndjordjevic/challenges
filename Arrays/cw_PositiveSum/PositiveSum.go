package main

import "fmt"

func main() {
	fmt.Println(PositiveSum([]int{1, -4, 7, 12}))
}

func PositiveSum(numbers []int) int {
	var s int

	for _, value := range numbers {
		if value > 0 {
			s += value
		}
	}

	return s
}
