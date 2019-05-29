package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toLowerCase("Nenad"))
}

func toLowerCase(s string) string {
	var str strings.Builder

	for _, e := range s {
		if e >= 'A' && e <= 'Z' {
			str.WriteRune(e + 32)
		} else {
			str.WriteRune(e)
		}
	}

	return str.String()
}
