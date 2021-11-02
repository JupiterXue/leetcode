package main

import "fmt"

func sumArray(numsPart []int) int {
	res := 0
	for _, v := range numsPart {
		res += v
	}
	return res
}

func minSubArrayLen(target int, nums []int) int {
	// 双指针法
	minLen := len(nums)
	if sumArray(nums) < target {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j <= len(nums); j++ {
			if sumArray(nums[i:j]) >= target {
				if minLen > j - i {
					minLen = j - i
				}
				break
			}
		}
	}
	return minLen
}

func main() {
	target := 7
	//nums := []int{2,3,1,2,4,3}
	nums := []int{1,1,1,7}
	fmt.Println(minSubArrayLen(target, nums))
}