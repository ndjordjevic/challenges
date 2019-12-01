package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(stringCompression("aabcccccaaadd"))
}

func stringCompression(str string) string {
	res := strings.Builder{}
	count := 0

	for i := 0; i < len(str); i++ {
		count++

		if i+1 == len(str) || str[i] != str[i+1] {
			res.WriteRune(rune(str[i]))
			res.WriteString(strconv.Itoa(count))
			count = 0
		}
	}

	if len(res.String()) < len(str) {
		return res.String()
	} else {
		return str
	}
}
