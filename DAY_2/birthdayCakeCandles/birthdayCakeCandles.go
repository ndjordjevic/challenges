package main

import "fmt"

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3, 3}))
}

func birthdayCakeCandles(ar []int32) int32 {
	m := make(map[int32]int32)
	max := ar[0]

	for _, v := range ar {
		if v > max {
			max = v
		}

		m[v] = m[v] + 1
	}

	return m[max]
}
