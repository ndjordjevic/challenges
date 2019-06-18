package main

import "fmt"

func main() {
	fmt.Println(DNAtoRNA("GCAT"))
}

func DNAtoRNA(dna string) string {
	r := []rune(dna)

	for i := range r {
		if r[i] == 'T' {
			r[i] = 'U'
		}
	}

	return string(r)
}
