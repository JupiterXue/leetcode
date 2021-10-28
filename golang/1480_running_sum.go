package main

import "fmt"

func runningSum(nums []int) []int {
	res := []int{}
	for i := range nums {
		count := 0
		for j := 0; j <= i; j++ {
			count += nums[j]
		}
		res = append(res, count)
	}
	return res
}

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(runningSum(nums))
}
