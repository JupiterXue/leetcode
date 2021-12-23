package main

import "fmt"

func missingNumber(nums []int) int {
	start := 0
	end := len(nums)
	for i := 0; i <= len(nums)/2; i++{
		if nums[i] == start {
			start++
		} else {
			return start
		}
		if nums[len(nums)-i-1] == end {
			end--
		} else {
			return end
		}
	}
	return start
}

func main() {
	//nums := []int{1}
	nums := []int{0,1,3}
	//nums := []int{0,1,2,3,4,5,6,7,9}
	fmt.Println(missingNumber(nums))
}
