package main

import "fmt"

func main() {
	fmt.Println(superReducedString("baaccbd"))
}

func superReducedString(s string) string {
	var r string

	var reduced bool
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			i++
			reduced = true
			continue
		}

		r += string(s[i])
	}

	if reduced {
		r = superReducedString(r)
	}

	if r == "" {
		r = "Empty String"
	}

	return r
}
