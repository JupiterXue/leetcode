package main

import "fmt"

//func search2(nums []int, target int) int {
//	low := 0
//	high := len(nums) - 1
//	for low <= high {
//		i := (high-low)/2 + low
//		middle := nums[i]
//		if middle == target {
//			return i
//		}
//		if target > middle {
//			low = i + 1
//		} else {
//			high = i - 1
//		}
//	}
//	return -1
//}

func search2(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		i := (low+high) / 2
		mid := nums[i]
		if mid == target {
			return i
		} else if mid > target {
			high = i - 1
		} else {
			low = i + 1
		}
	}
	return -1
}


func main() {
	nums := []int{-1,0,3,5,9,12}
	//target := 2
	target := 2
	fmt.Println(search2(nums, target))
}
