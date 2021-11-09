package main

import (
	"fmt"
	"math"
)

func sumArray(numsPart []int) int {
	res := 0
	for _, v := range numsPart {
		res += v
	}
	return res
}

//func minSubArrayLen(target int, nums []int) int {
//	// 双指针法
//	minLen := len(nums)
//	if sumArray(nums) < target {
//		return 0
//	}
//	for i := 0; i < len(nums); i++ {
//		for j := i; j <= len(nums); j++ {
//			if sumArray(nums[i:j]) >= target {
//				if minLen > j - i {
//					minLen = j - i
//				}
//				break
//			}
//		}
//	}
//	return minLen
//}



//func minSubArrayLen(target int, nums []int) int {
//	low := 0
//	high := 0
//	sum := 0
//	INTMAX := int(math.Pow(2.0,31) - 1)
//	min := INTMAX
//	for high < len(nums) {
//		sum += nums[high]
//		high++
//		for sum >= target {
//			min = min_num(min, high - low)
//			sum -= nums[low]
//			low++
//		}
//	}
//	if min == INTMAX {
//		return 0
//	} else {
//		return min
//	}
//}

func min_num(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minSubArrayLen(target int, nums []int) int {
	low, high, sum, min := 0, 0, 0, math.MaxInt32
	for high < len(nums) {
		sum += nums[high]
		high++
		for sum >= target {
			min = min_num(min, high- low)
			sum -= nums[low]
			low++
		}
	}
	if min == math.MaxInt32 {
		return 0
	} else {
		return min
	}
}



func main() {
	target := 7
	//nums := []int{2,3,1,2,4,3}
	nums := []int{1,1,1,7}
	fmt.Println(minSubArrayLen(target, nums))
}