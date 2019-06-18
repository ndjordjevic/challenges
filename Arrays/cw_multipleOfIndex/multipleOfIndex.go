package main

import (
	"fmt"
)

func main() {
	fmt.Println(multipleOfIndex([]int{0, -6, 32, 82, 9, 25}))
}

func multipleOfIndex(ints []int) []int {
	var r []int

	for key, value := range ints {
		if key == 0 {
			if value == 0 {
				r = append(r, value)
			}
			continue
		}

		if value%key == 0 {
			r = append(r, value)
		}
	}

	return r
}
