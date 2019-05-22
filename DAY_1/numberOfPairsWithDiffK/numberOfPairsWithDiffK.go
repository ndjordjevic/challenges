package numberOfPairsWithDiffK

import (
	"strconv"
)

func numberOfPairsWithDiffK(a []int, k int) int {
	//var n = 0

	//for i:=0; i < len(a); i++ {
	//	for j:=i+1; j < len(a); j++ {
	//		if a[i] + k == a[j] || a[i] - k == a[j] {
	//			n++
	//			fmt.Println(a[i], a[j])
	//		}
	//	}
	//}

	//sort.Ints(a)
	//for _, e := range a {
	//	i := sort.Search(len(a), func(i int) bool { return a[i] >= e + k || a[i] >= e - k})
	//
	//	if i < len(a) && (a[i] == e + k || a[i] == e - k) {
	//		n++
	//		fmt.Println(e, a[i])
	//	}
	//}

	m := make(map[int]bool)

	for _, e := range a {
		m[e] = true
	}

	set := make(map[string]bool)

	for _, e := range a {
		if m[e+k] {
			if e < e+k {
				set[strconv.Itoa(e)+" + "+strconv.Itoa(e+k)] = true
			} else {
				set[strconv.Itoa(e+k)+" + "+strconv.Itoa(e)] = true
			}
		}

		if m[e-k] {
			if e < e-k {
				set[strconv.Itoa(e)+" + "+strconv.Itoa(e-k)] = true
			} else {
				set[strconv.Itoa(e-k)+" + "+strconv.Itoa(e)] = true
			}
		}
	}

	return len(set)
}
