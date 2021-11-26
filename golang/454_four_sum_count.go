package main

import (
	"fmt"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	res := 0
	count := map[int]int{}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			count[v1+v2]++
		}
	}
	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			res += count[-v1-v2]
		}
	}
	return res
}

func main() {
	nums1, nums2, nums3, nums4 := []int{1,2}, []int{-2,-1},[]int{-1,2},[]int{0,2}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
