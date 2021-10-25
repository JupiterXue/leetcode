package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

func main() {
	a := 123
	b := 16
	fmt.Println(getSum(a, b))
}
