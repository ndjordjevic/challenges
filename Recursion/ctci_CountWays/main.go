package main

import "fmt"

func main() {
	fmt.Println(countWays2(5))
}

func countWays(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		return countWays(n-1) + countWays(n-2) + countWays(n-3)
	}
}

func countWays2(n int) int {
	memo := make([]int, n+1)

	for i, _ := range memo {
		memo[i] = -1
	}

	return countWaysRec(n, memo)
}

func countWaysRec(n int, memo []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		memo[n] = countWaysRec(n-1, memo) + countWaysRec(n-2, memo) + countWaysRec(n-3, memo)
		return memo[n]
	}
}
