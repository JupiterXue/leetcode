package main

import "fmt"

// 1. 自己做，动态，failed
//func canJump(nums []int) bool {
//	if len(nums) == 1 {
//		return true
//	}
//	last := len(nums)-1
//	for i := 0; i < len(nums); {
//		if nums[i] == 0 {
//			return false
//		}
//		fmt.Println(nums[i])
//		if nums[i+1] > nums[i] {
//			i += nums[i+1]
//		} else {
//			i += nums[i]
//		}
//		fmt.Println(nums[i])
//		if i >= last {
//			return true
//		}
//	}
//	return false
//}

// 2. 题解，随想录，动态
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i:= 1; i < len(nums); i++ {
		for j := i-1; j >= 0; j-- {
			if dp[j] && nums[j] + j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}


func main() {
	//nums := []int{2,3,1,1,4}
	//nums := []int{3,2,1,0,4}
	nums := []int{1,2,0,1}
	fmt.Println(canJump(nums))
}