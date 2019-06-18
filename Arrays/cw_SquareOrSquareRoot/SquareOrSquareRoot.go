package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}))
}

func SquareOrSquareRoot(arr []int) []int {
	r := make([]int, len(arr))

	for i, v := range arr {
		sqrt := math.Sqrt(float64(v))
		if sqrt == math.Floor(sqrt) {
			r[i] = int(sqrt)
		} else {
			r[i] = int(math.Pow(float64(v), float64(2)))
		}
	}

	return r
}
