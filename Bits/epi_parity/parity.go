package main

import "fmt"

func main() {
	fmt.Println("\nParity:", parity(100))
}
func parity(x uint64) uint64 {
	var res uint64

	for x > 0 {
		fmt.Printf("%v, %b, %v, %v\n", x, x, res, x&1)
		res ^= (x & 1)
		x >>= 1
	}

	return res
}
