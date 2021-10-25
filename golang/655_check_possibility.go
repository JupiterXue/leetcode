package main

import (
	"fmt"
)


func checkPossibility(nums []int) bool {
	lenNums := len(nums)
	count := 0
	for i := 1; i < lenNums ; i++ {
		if nums[i] < nums[i-1] {
			count++
			if i== 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return count <= 1
}

func main() {
	//nums := []int{4,2,3}
	//nums := []int{4,2,1}
	nums := []int{3,4,2,3}
	fmt.Println(checkPossibility(nums))
}
