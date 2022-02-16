package main

import (
	"fmt"
)

// 1. 自己写，Overtime 复杂度过高， O(n^2)
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

// 2. 题解，AC
//func maxSubArray(nums []int) int {
//	sum := nums[0]
//	maxValue := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if sum < 0 {
//			sum = nums[i]
//		} else {
//			sum += nums[i]
//		}
//		if maxValue < sum {
//			maxValue = sum
//		}
//	}
//	return maxValue
//}



// 3. 自己写，动态规划，没写出来。
//func maxSubArray(nums []int) int {
//	max := nums[0]
//	for i := 0; i < len(nums); i++ {
//		for j := i; j < len(nums); j++ {
//			curSum := sum(nums[j:])
//			fmt.Println(nums[j:])
//			if max < curSum {
//				max = curSum
//			} else {
//				break
//			}
//		}
//	}
//	return max
//}

//func sum(s []int) int {
//	count := 0
//	for _, v := range s {
//		count += v
//	}
//	return count
//}

// 4. 题解，动态规划，AC！
//func maxSubArray(nums []int) int {
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] + nums[i-1] > nums[i] {
//			nums[i] += nums[i-1]
//		}
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	return max
//}

// 5. 自己写，贪婪，AC
func maxSubArray(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > sum {
			sum = nums[i]
		}
	}
	return sum
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	//nums := []int{5,4,-1,7,8}
	fmt.Println(maxSubArray(nums))
}
