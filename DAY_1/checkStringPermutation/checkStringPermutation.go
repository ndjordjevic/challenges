package main

import "fmt"

func main() {
	fmt.Println(checkStringPermutation("nenaddj", "naddnej"))
}

//func checkStringPermutation(s1 string, s2 string) bool {
//	var m1 = make(map[int32] int)
//	var m2 = make(map[int32] int)
//
//	for _, e := range s1 {
//		m1[e] = m1[e] + 1
//	}
//
//	for _, e := range s2 {
//		m2[e] = m2[e] + 1
//	}
//
//	for _, e := range s1 {
//		if m1[e] != m2[e] {
//			return false
//		}
//	}
//
//	return true
//}

func checkStringPermutation(s1 string, s2 string) bool {
	var a [128]int

	for _, v := range s1 {
		a[v]++
	}

	for _, v := range s2 {
		a[v]--
		if a[v] < 0 {
			return false
		}
	}

	return true
}
