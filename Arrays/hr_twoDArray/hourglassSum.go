package main

import "fmt"

func main() {
	fmt.Println(hourglassSum([][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}))
}

func hourglassSum(arr [][]int32) int32 {
	var max int32
	var s int32

	max = -63
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if s > max {
				max = s
			}
		}
	}

	return max
}
