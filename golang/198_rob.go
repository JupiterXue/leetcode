package main

import "fmt"
func max(n []int) int {
	maxValue := n[0]
	for _, v := range n {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func maxNum(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) <= 2 {
		return max(nums)
	}
	dp := []int{nums[0], maxNum(nums[0], nums[1])}
	for i, _ := range nums {
		if i < 2 {
			continue
		}
		dp = append(dp, maxNum(dp[i-2] + nums[i], dp[i-1]))
	}
	return max(dp)
}

func main() {
	nums := []int{1,2,3,1}
	fmt.Println(rob(nums))
}
