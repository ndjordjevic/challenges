package sortArrayByParity

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	var r []int

	for _, e := range A {
		if e%2 == 0 {
			// insert at the begging (prepend)
			r = append([]int{e}, r...)
		} else {
			r = append(r, e)
		}
	}

	return r
}
