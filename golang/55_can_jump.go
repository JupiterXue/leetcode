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
//func canJump(nums []int) bool {
//	if len(nums) <= 1 {
//		return true
//	}
//	dp := make([]bool, len(nums))
//	dp[0] = true
//	for i:= 1; i < len(nums); i++ {
//		for j := i-1; j >= 0; j-- {
//			if dp[j] && nums[j] + j >= i {
//				dp[i] = true
//				break
//			}
//		}
//	}
//	return dp[len(nums)-1]
//}

// 3. 自己写，贪婪，Failed 5,9,3,2,1,0,2,3,3,1,0,0
//func canJump(nums []int) bool {
//	if len(nums) < 1 {
//		return true
//	}
//	cur, step := nums[0], nums[0]
//	for i := 1; i < len(nums); i++ {
//		if step == 0 {
//			return false
//		}
//		if nums[i] >= cur - 1 {
//			cur = nums[i]
//			step = nums[i]
//		} else {
//			step--
//		}
//	}
//	return true
//}


// 4. 随想录，贪婪，Faild...
//func canJump(nums []int) bool {
//	if len(nums) <= 1 {
//		return true
//	}
//	cover := 0
//	for i := 0; i <= len(nums); i++ {
//		if i + nums[i] > cover {
//			cover = i + nums[i]
//		}
//		if cover >= len(nums) - 1 {
//			return true
//		}
//	}
//	return false
//}

// 5. 题解
func canJump(nums []int) bool {
	length := len(nums) - 1
	for i := length-1; i >= 0; i-- {
		fmt.Println(i)
		fmt.Println(nums[i] + i)
		fmt.Println()
		if nums[i] + i >= length {
			length = i
		}
	}
	return length <= 0
}

func main() {
	nums := []int{2,3,1,1,4}
	//nums := []int{3,2,1,0,4}
	//nums := []int{1,2,0,1}
	//nums := []int{1,1,2,2,0,1,1}
	//nums := []int{5,9,3,2,1,0,2,3,3,1,0,0}
	fmt.Println(canJump(nums))
}