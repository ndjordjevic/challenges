package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(AbbrevName("Sam Harris"))
}

func AbbrevName(name string) string {
	//return string(unicode.ToUpper(rune(name[0]))) + "." + string(unicode.ToUpper(rune(name[strings.Index(name, " ")+1])))

	words := strings.Split(name, " ")
	return strings.ToUpper(string(words[0][0])) + "." + strings.ToUpper(string(words[1][0]))
}
