package main

import "fmt"

func main() {
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	fmt.Printf("%s", s)
}

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l <= r; l, r = l+1, r-1 {
		t := s[l]
		s[l] = s[r]
		s[r] = t
	}
}
