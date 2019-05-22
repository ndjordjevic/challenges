package main

import "fmt"

func main() {
	staircase(8)
}

func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		for j := i + 1; j < int(n); j++ {
			fmt.Printf(" ")
		}

		for k := int(n) - i - 1; k < int(n); k++ {
			fmt.Printf("#")
		}

		fmt.Println()
	}
}
