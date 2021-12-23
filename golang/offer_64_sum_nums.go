package main

import "fmt"

// 1. 题解
func sumNums(n int) int {
	res := 0
	var sum func(int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return res
}

func main() {
	n := 3
	fmt.Println(sumNums(n))
}
