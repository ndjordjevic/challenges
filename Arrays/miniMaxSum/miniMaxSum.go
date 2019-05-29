package main

import (
	"fmt"
	"sort"
)

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
}

func miniMaxSum(arr []int32) {

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var min, max int64

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			min += int64(arr[i])
		} else if i == len(arr)-1 {
			max += int64(arr[i])
		} else {
			min += int64(arr[i])
			max += int64(arr[i])
		}
	}

	fmt.Println(min, max)
}
