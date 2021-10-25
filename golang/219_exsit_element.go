package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums) - 1; i++ {
		for j := i+1; j < len(nums); j++ {
			if j - i > k {
				break
			}
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{1,0,1,1}
	k := 1
	fmt.Println(containsNearbyDuplicate(nums, k))
}
