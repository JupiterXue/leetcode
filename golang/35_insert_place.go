package main

import "fmt"

//func searchInsert(nums []int, target int) int {
//	for i, v := range nums {
//		if v == target {
//			return i
//		}
//	}
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > target {
//			return i
//		}
//	}
//	return len(nums)
//}

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		//i = (high-low) / 2 + low
		i := (high+low) / 2
		mid := nums[i]
		if mid == target {
			return i
		} else if mid > target {
			high = i - 1
		} else {
			low = i + 1
		}
	}
	return low
}

func main() {
	//nums := []int{1,3,5,6}
	nums := []int{1,3}
	target := 2
	fmt.Println(searchInsert(nums, target))
}
