package main

import (
	"fmt"
	"math"
)



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

//func sumArray(numsPart []int) int {
//	res := 0
//	for _, v := range numsPart {
//		res += v
//	}
//	return res
//}
//
//func min_num(a, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}
//
//func minSubArrayLen(target int, nums []int) int {
//	low, high, sum, min := 0, 0, 0, math.MaxInt32
//	for high < len(nums) {
//		sum += nums[high]
//		high++
//		for sum >= target {
//			min = min_num(min, high- low)
//			sum -= nums[low]
//			low++
//		}
//	}
//	if min == math.MaxInt32 {
//		return 0
//	} else {
//		return min
//	}
//}


// 2. 滑动窗口，AC！
func minSubArrayLen(target int, nums []int) int {
	// 窗口大小，窗口位置，窗口左边界
	size, sum, index := math.MaxInt32, 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 当窗口内的和大于目标值是，左边界右移，缩小边界。
		for sum >= target {
			// 取更小的窗口，i - index + 1 是窗口长度，即子数组长度
			if size < i - index + 1 {
				size = size
			} else {
				size = i - index + 1
			}
			// 窗口左边界右移
			sum -= nums[index]
			index++
		}
	}
	// 判断窗口大小是否变化，如果变化说明找到了
	if size != math.MaxInt32 {
		return size
	} else {
		return 0
	}
}


func main() {
	target := 7
	//nums := []int{2,3,1,2,4,3}
	nums := []int{1,1,1,7}
	fmt.Println(minSubArrayLen(target, nums))
}