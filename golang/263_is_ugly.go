package main

import "fmt"

func isUgly(n int) bool {
	if n < 1 {
		return false
	}

	uglyList := []int{2,3,5}
	for _, v := range uglyList {
		for n % v == 0 {
			n /= v
		}
	}

	return n == 1
}

func main() {
	n := 6
	fmt.Println(isUgly(n))
}
