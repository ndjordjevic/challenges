package main

import "fmt"

func main() {
	fmt.Println(fibonacci3(1))
}

func fibonacci(i int) int {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return fibonacci(i-1) + fibonacci(i-2)
}

func fibonacci2(n int) int {
	a := make([]int, n+1)
	return fibonacciRec(n, a)
}

func fibonacciRec(i int, memo []int) int {
	if i == 0 || i == 1 {
		return i
	}

	if memo[i] == 0 {
		memo[i] = fibonacciRec(i-1, memo) + fibonacciRec(i-2, memo)
	}

	return memo[i]
}

func fibonacci3(n int) int {
	if n == 0 {
		return 0
	}

	a, b := 0, 1
	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}
	return a + b
}
