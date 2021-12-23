package main

import "fmt"

// 1. 自己写，超时
//func getDescentPeriods(prices []int) int64 {
//	size := 1
//	res := 0
//	for size <= len(prices) {
//		for i := 0; i <= len(prices) - size; i++ {
//			window := prices[i:i+size]
//			if isDescent(window) {
//				fmt.Println(window)
//				res++
//			}
//		}
//		size++
//	}
//	return int64(res)
//}
//
//func isDescent(nums []int) bool {
//	for i := 1; i < len(nums); i++ {
//		if nums[i] - nums[i-1] != -1  {
//			return false
//		}
//	}
//	return true
//}

// 2. 自己写，DP
func getDescentPeriods(prices []int) int64 {
	lastIdx := 0
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] - prices[i-1] != -1 {
			res += count(i-lastIdx)
			fmt.Println(i-lastIdx)
			fmt.Println(res)
			lastIdx = i+1
			fmt.Println()
		}
	}
	return int64(res)
}

func count(n int) int {
	if n == 1 {
		return 1
	}
	return n + count(n)
}

func main() {
	prices := []int{3,2,1,4}
	//prices := []int{8,6,7,7}
	//prices := []int{1}
	fmt.Println(getDescentPeriods(prices))
}
