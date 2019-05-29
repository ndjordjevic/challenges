package main

import "fmt"

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}

func plusMinus(arr []int32) {
	l := len(arr)
	var pos, zero, neg int

	for _, v := range arr {
		if v < 0 {
			neg++
		} else if v > 0 {
			pos++
		} else {
			zero++
		}
	}

	fmt.Printf("%.6f\n", float32(pos)/float32(l))
	fmt.Printf("%.6f\n", float32(neg)/float32(l))
	fmt.Printf("%.6f\n", float32(zero)/float32(l))

}
