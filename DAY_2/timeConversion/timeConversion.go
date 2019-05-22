package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(timeConversion("06:40:03AM"))
}

func timeConversion(s string) string {
	hh := s[0:2]
	ap := s[8:10]
	mm := s[3:5]
	ss := s[6:8]
	var newHHStr string

	if hh == "12" && ap == "PM" {
		newHHStr = hh
	}

	if hh == "12" && ap == "AM" {
		newHHStr = "00"
	}

	hhN, _ := strconv.Atoi(hh)
	if hhN < 12 && ap == "AM" {
		newHHStr = "0" + strconv.Itoa(hhN)
	}

	if hhN < 12 && ap == "PM" {
		newHHStr = strconv.Itoa(hhN + 12)
	}

	return newHHStr + ":" + mm + ":" + ss
}
