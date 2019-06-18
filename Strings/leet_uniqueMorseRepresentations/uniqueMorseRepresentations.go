package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

func uniqueMorseRepresentations(words []string) int {
	t := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	var m = make(map[string]bool)
	var s strings.Builder

	for i := range words {
		s.Reset()
		for j, _ := range words[i] {
			s.WriteString(t[words[i][j]-'a'])
		}
		m[s.String()] = true
	}

	return len(m)

}
