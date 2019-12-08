package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(math.MaxInt32)
	n, _ := strconv.ParseInt("10000000", 2, 32)
	n32 := int32(n)
	fmt.Println(n32)
	m, _ := strconv.ParseInt("10011", 2, 32)
	m32 := int32(m)
	fmt.Println(m32)

	fmt.Printf("%b\n", updateBits(n32, m32, 1, 5))
}

func updateBits(n, m, i, j int32) int32 {
	all1, err := strconv.ParseInt("1111111111111111111111111111111", 2, 32)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", all1)
	allOnes := int32(^0)
	all1u := uint32(all1)
	fmt.Printf("%v\n", all1u)
	fmt.Printf("%b\n", ^5)

	left := allOnes << (j + 1)
	fmt.Printf("Left: %b\n", left)
	right := int32((2 << i) - 1)
	mask := int32(left | right)

	nCleared := n & mask
	mShifted := m << i

	return nCleared | mShifted
}
