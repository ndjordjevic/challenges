package main

import (
	"fmt"
)

func main() {
	fmt.Println(intersection2([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]interface{})
	var r []int
	for _, v1 := range nums1 {
		if _, ok := m[v1]; ok {
			continue
		}
		for _, v2 := range nums2 {
			if _, ok := m[v2]; v1 == v2 && !ok {
				m[v1] = nil
				r = append(r, v1)
			}
		}
	}

	return r
}

func intersection2(nums1 []int, nums2 []int) []int {
	var count = map[int]bool{}
	var result []int
	for _, num := range nums1 {
		count[num] = true
	}
	for _, num := range nums2 {
		if count[num] == true {
			result = append(result, num)
			count[num] = false
		}
	}
	return result
}
