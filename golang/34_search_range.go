package main

import "fmt"

//func searchRange(nums []int, target int) []int {
//	res := []int{-1, -1}
//	for i, v := range nums {
//		if v == target {
//			if res[0] == -1 {
//				res[0] = i
//			} else {
//				res[1] = i
//			}
//		}
//	}
//	if res[0] != -1 && res[1] == -1 {
//		res[1] = res[0]
//	}
//	return res
//}

func searchRange(nums []int, target int) []int {
	// 双指针
	res := []int{-1,-1}
	for i := 0; i < len(nums); i++ {
		if res[0] == -1 {
			if nums[i] == target {
				res[0] = i
			}
		}
		if res[1] == -1 {
			if nums[len(nums)-i-1] == target {
				res[1] = len(nums) - i
			}
		}
	}
	return res
}


func main() {
	nums := []int{5,7,7,8,8,10}
	target := 8
	fmt.Println(searchRange(nums, target))
}