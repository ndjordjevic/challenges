package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
}

// Complete the camelcase function below.
func camelcase(s string) int32 {
	var r int32

	for _, v := range s {
		if unicode.IsUpper(v) {
			r++
		}
	}

	return r + 1
}
