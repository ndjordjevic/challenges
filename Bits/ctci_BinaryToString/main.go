package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printBinary(0.56))
}

func printBinary(num float32) string {
	if num >= 1 || num <= 0 {
		return "ERROR"
	}

	b := strings.Builder{}
	b.WriteRune('.')
	for num > 0 {
		if b.Len() > 32 {
			return "ERROR"
		}

		r := num * 2
		if r >= 1 {
			b.WriteRune('1')
			num = r - 1
		} else {
			b.WriteRune('0')
			num = r
		}
	}

	return b.String()
}
