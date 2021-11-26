package main

import "fmt"

// 1. 暴力法
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

// 2. 二分法
//func searchInsert(nums []int, target int) int {
//	low := 0
//	high := len(nums) - 1
//	for low <= high {
//		//i = (high-low) / 2 + low
//		i := (high+low) / 2
//		mid := nums[i]
//		if mid == target {
//			return i
//		} else if mid > target {
//			high = i - 1
//		} else {
//			low = i + 1
//		}
//	}
//	return low
//}

// 3. 二分递归法
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	return recursionBinarySearch(nums, low, high, target)
}

func recursionBinarySearch(nums []int, low, high, target int) int {
	if low > high {
		return low
	}
	mid :=(low + high) / 2
	if nums[mid] == mid {
		return mid
	} else if nums[mid] < target {
		return recursionBinarySearch(nums, mid+1, high, target)
	} else {
		return recursionBinarySearch(nums, low, mid-1, target)
	}
}


func main() {
	nums := []int{1,3,5,6}
	//nums := []int{1,3}
	target := 0
	fmt.Println(searchInsert(nums, target))
}
