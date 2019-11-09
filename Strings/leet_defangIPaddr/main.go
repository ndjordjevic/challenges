package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr2("255.100.50.0"))
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func defangIPaddr2(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}
