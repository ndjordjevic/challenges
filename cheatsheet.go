package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// declare and instance a map
	var m1 = make(map[int32]int)
	fmt.Println(m1)

	//iterate over a string, array, slice, map
	for i, v := range "sdfsd" {
		fmt.Println(i, v)
	}

	// declare an array
	var a [128]int
	fmt.Println(a)

	// declare a slice with a specific length
	r := make([]int32, 2)
	fmt.Println(r)

	// declare multiple vars in one line of the same type
	var s1, s2 int32
	fmt.Println(s1, s2)

	// length of array, slice, map, string (chars but not runes)
	var arr [5]int
	fmt.Println(len(arr))

	// multiple (i, j) vars in for loop
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		//s1 += arr[i][i]
		//s2 += arr[i][j]
	}

	// Math.Abs
	fmt.Println(math.Abs(-44.3))

	// two dimension slice literal
	td := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(td)

	// sort an int array
	sort.Ints([]int{4, 5, 6})

	// binary search a sorted slice
	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= 3 })
	fmt.Println(i)

	// declare a set like map
	set := make(map[string]bool)
	fmt.Println(set)

	// convert a int to string
	fmt.Println(strconv.Itoa(5))

	// convert a string to int
	fmt.Println(strconv.Atoi("5"))

	// sort a slice
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// map of rune and int
	m := make(map[rune]int)
	fmt.Println(m)

	// lower case a char/rune
	fmt.Println(unicode.ToLower('F'))

	// a slice of runes from a string
	b := []rune("Nenad")
	fmt.Println(b)

	// is rune in string upper, IsDigit, IsLower ... the similar functions
	fmt.Println(unicode.IsUpper('R'))

	// check if string contains any of provided characters
	strings.ContainsAny("sdfsd#", "!@#$%^&*()-+")

	// string Builder, WriteString or WriteRune
	var strB strings.Builder
	strB.WriteString("Hello")
	strB.WriteString(" World")
	strB.Reset() // reset a String Builder
}
