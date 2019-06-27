package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}

func flipAndInvertImage(A [][]int) [][]int {
	l := len(A)
	var t int

	for i := 0; i < l; i++ {
		for j := 0; j < len(A)/2; j++ {
			t = A[i][j]
			A[i][j] = A[i][l-j-1]
			A[i][l-j-1] = t
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			A[i][j] = 1 - A[i][j]
		}
	}

	return A
}
