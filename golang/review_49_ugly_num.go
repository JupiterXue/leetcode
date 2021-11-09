package main

import "fmt"

func minNums(nums []int) int {
	minNum := nums[0]
	for _, v := range nums {
		if minNum < v {
			minNum = v
		}
	}
	return minNum
}

func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[1] = 1
	a2, a3, a5:= 1, 1 ,1
	for i := 2; i <=n; i++ {
		min := dp[a2]*2
		if dp[a3]*3 < min {
			min = dp[a3]*3
		}
		if dp[a5]*5 < min {
			min = dp[a5]*5
		}
		dp[i] = min

		if min == dp[a2]*2{
			a2++
		}
		if min == dp[a3]*3{
			a3++
		}
		if min == dp[a5]*5{
			a5++
		}
	}
	return dp[n]
}

func main() {
	n := 10
	fmt.Println(nthUglyNumber(n))
}