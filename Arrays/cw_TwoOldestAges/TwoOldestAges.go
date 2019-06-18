package main

import "fmt"

func main() {
	fmt.Println(TwoOldestAges([]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}))
}

func TwoOldestAges(ages []int) [2]int {
	var r [2]int

	for _, value := range ages {
		if value > r[1] {
			r[0] = r[1]
			r[1] = value
			continue
		}

		if value > r[0] {
			r[0] = value
		}
	}

	return r
}
