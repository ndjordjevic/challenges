package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes3([]int{2, 1, 2, 5, 3, 2}))
}

func repeatedNTimes(A []int) int {
	n := len(A) / 2
	m := make(map[int]int)

	for _, v := range A {
		m[v] = m[v] + 1
		if m[v] == n {
			return v
		}
	}

	return -1
}

func repeatedNTimes2(A []int) int {
	var mp [10000]int
	l := len(A)
	for i := 0; i < l; i++ {
		mp[A[i]]++
		if mp[A[i]] > 1 {
			return A[i]
		}
	}
	return 0
}

func repeatedNTimes3(A []int) int {
	seen := make(map[int]bool)

	for _, num := range A {
		if seen[num] {
			return num
		} else {
			seen[num] = true
		}
	}
	return -1
}
