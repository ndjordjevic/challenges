package main

import (
	"fmt"
)

func main() {
	m := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	rotateMatrix(m)

	fmt.Println(m)
}

func rotateMatrix(m [][]int) {
	N := len(m)
	if N == 0 || N != len(m[0]) {
		return
	}

	for i := 0; i < N/2; i++ {
		for j := i; j < N-i-1; j++ {
			temp := m[i][j]
			m[i][j] = m[N-1-j][i]
			m[N-1-j][i] = m[N-1-i][N-1-j]
			m[N-1-i][N-1-j] = m[j][N-1-i]
			m[j][N-1-i] = temp
		}
	}
}
