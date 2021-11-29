package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	res := []int{}
	for i, v := range nums {
		if v == target {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	nums := []int{1,2,5,2,3}
	target := 4
	fmt.Println(targetIndices(nums, target))
}
