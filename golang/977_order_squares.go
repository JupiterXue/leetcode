package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		res = append(res, v*v)
	}
	sort.Ints(res)
	return res
}

func main() {
	nums := []int{-4,-1,0,3,10}
	fmt.Println(sortedSquares(nums))
}
