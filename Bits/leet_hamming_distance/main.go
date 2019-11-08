package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
