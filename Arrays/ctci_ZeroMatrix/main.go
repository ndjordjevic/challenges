package main

import "fmt"

func main() {
	matrix := [][]int{{0, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	setZeros(matrix)

	fmt.Println(matrix)
}

func setZeros(matrix [][]int) {
	row := make([]bool, len(matrix))
	column := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				column[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		if row[i] {
			nullifyRow(matrix, i)
		}
	}

	for i := 0; i < len(column); i++ {
		if column[i] {
			nullifyColumn(matrix, i)
		}
	}
}

func nullifyRow(matrix [][]int, row int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[row][i] = 0
	}
}

func nullifyColumn(matrix [][]int, column int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][column] = 0
	}
}
