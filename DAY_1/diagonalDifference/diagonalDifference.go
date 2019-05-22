package main

import (
	"fmt"
	"math"
)

//func diagonalDifference(arr [][]int32) int32 {
//	var s1, s2 int32
//
//	for i := 0; i < len(arr); i++ {
//		for j := 0; j < len(arr); j++ {
//			if i == j {
//				s1 += arr[i][j]
//			}
//
//			if (i + j) == len(arr)-1 {
//				s2 += arr[i][j]
//			}
//		}
//	}
//
//	return int32(math.Abs(float64(s1 - s2)))
//}

func diagonalDifference(arr [][]int32) int32 {
	var s1, s2 int32

	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		s1 += arr[i][i]
		s2 += arr[i][j]
	}

	return int32(math.Abs(float64(s1 - s2)))
}

func main() {
	fmt.Print(diagonalDifference([][]int32{
		{11, 2, 4}, {4, 5, 6}, {10, 8, -12},
	}))
}
