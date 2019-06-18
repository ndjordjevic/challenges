package main

import "fmt"

func main() {
	fmt.Println(FindUniq([]float32{1, 1, 1, 2, 1, 1}))
}

func FindUniq(arr []float32) float32 {
	m := make(map[float32]float32)

	for _, value := range arr {
		m[value] += 1
	}

	for key, value := range m {
		if value == 1 {
			return key
		}
	}

	return 0
}
