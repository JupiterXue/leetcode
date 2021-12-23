package main

import (
	"fmt"
	"sort"
)

// 1. 自己写，failed，吐了，周赛还做通过，现在居然不会了。
//func maxSubsequence(nums []int, k int) []int {
//	deleteN := len(nums)-k
//	for i := 0; i <= deleteN; i++ {
//		minIdx := getMinIdx(nums)
//		nums = append(nums[:minIdx], nums[minIdx+1:]...)
//	}
//	return nums
//}
//
//func getMinIdx(nums []int) int {
//	min := nums[0]
//	minIdx := 0
//	for i, v := range nums {
//		if v <= min {
//			min = v
//			minIdx = i
//		}
//	}
//	return minIdx
//}

// 2. 题解，AC
func maxSubsequence(nums []int, k int) []int {
	id := make([]int, len(nums))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return nums[id[i]] > nums[id[j]]
	})
	sort.Ints(id[:k])
	res := make([]int, k)
	for i, j := range id[:k] {
		res[i] = nums[j]
	}
	return res
}


func main() {
	//nums := []int{2,1,3,3}
	//nums := []int{-1,-2,3,4}
	//nums := []int{3,4,3,3}
	//nums := []int{50,-75}
	nums := []int{33,-27,-9,-83,48}
	//k := 2
	k := 3
	fmt.Println(maxSubsequence(nums, k))
}
