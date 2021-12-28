package main

import (
	"fmt"
)

// 1. 自己写，去重+快排+比较大小，超时
//func longestConsecutive(nums []int) int {
//	quicksort(nums, 0, len(nums)-1)
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//		}
//	}
//	fmt.Println(nums)
//	max := len(nums)
//	//fmt.Println(max)
//	for max > 1 {
//		for i := 0; i <= len(nums) - max; i++ {
//			fmt.Println(nums[i:i+max])
//			if checkConsecutive(nums[i:i+max]) {
//				return max
//			}
//		}
//		max--
//	}
//	return max
//}
//
//func checkConsecutive(nums []int) bool {
//	for i := 0; i < len(nums)-1; i++ {
//		if getSub(nums[i],nums[i+1]) > 1 {
//			return false
//		}
//	}
//	return true
//}
//
//func getSub(a, b int) int {
//	if a > b {
//		return a - b
//	} else {
//		return b - a
//	}
//}
//
//func quicksort(nums []int, low int, high int) {
//	if low < high {
//		pivotIdx := newPartition(nums, low, high)
//		quicksort(nums, low, pivotIdx-1)
//		quicksort(nums, pivotIdx+1, high)
//	}
//}
//
//func newPartition(nums []int, low int, high int) int {
//	pivot := nums[low]
//	for low < high {
//		for low < high && pivot <= nums[high] {
//			high--
//		}
//		nums[low] = nums[high]
//		for low < high && pivot >= nums[low] {
//			low++
//		}
//		nums[high] = nums[low]
//	}
//	nums[low] = pivot
//	return low
//}


// 2. 题解，AC！
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	//fmt.Println(numSet)
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			cur := num
			curStreak := 1
			for numSet[cur+1] {
				cur++
				curStreak++
			}
			if longestStreak < curStreak {
				longestStreak = curStreak
			}
		}
	}
	return longestStreak
}

func main() {
	//nums := []int{100,4,200,1,3,2}
	//nums := []int{1,2,0,1}
	nums := []int{1,-8,7,-2,-4,-4,6,3,-4,0,-7,-1,5,1,-9,-3}
	fmt.Println(longestConsecutive(nums))
}
