package main

import (
	"fmt"
)

// 1. 直接遍历+排序，AC！
//func sortedSquares(nums []int) []int {
//	res := []int{}
//	for _, v := range nums {
//		res = append(res, v*v)
//	}
//	sort.Ints(res)
//	return res
//}

// 2. 双指针
func sortedSquares(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	low, high := 0, length-1
	for pos := length-1; pos >= 0; pos-- {
		if v1, v2 := nums[low] * nums[low], nums[high] * nums[high]; v1 > v2 {
			res[pos] = v1
			low++
		} else {
			res[pos] = v2
			high--
		}
	}
	return res
}

func main() {
	nums := []int{-4,-1,0,3,10}
	fmt.Println(sortedSquares(nums))
}
