package main

import (
	"fmt"
)

func sum(s []int) int {
	count := 0
	for _, v := range s {
		count += v
	}
	return count
}

// 复杂度过高， O(n^2)
//func maxSubArray(nums []int) int {
//	maxValue := nums[0]
//	for i := 0; i < len(nums); i++ {
//		for j := i; j < len(nums); j++ {
//			value := sum(nums[i:j+1])
//			if value > maxValue {
//				maxValue = value
//			}
//		}
//	}
//	return maxValue
//}

func maxSubArray(nums []int) int {
	sum := nums[0]
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if maxValue < sum {
			maxValue = sum
		}
	}
	return maxValue
}

func main() {
	nums := []int{5,4,-1,7,8}
	fmt.Println(maxSubArray(nums))
}
