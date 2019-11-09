package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))
}

func judgeCircle(moves string) bool {
	var x, y int

	for _, m := range moves {
		switch m {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++

		}
	}

	return x == 0 && y == 0
}
