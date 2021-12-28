package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 3 {
		for i := 1; i < len(nums); i++ {
			if nums[i] == nums[i-1] {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		}
	}
	for i := 1; i < len(nums) - 1; i ++ {
		if (nums[i+1] - nums[i] > 0 && nums[i] - nums[i-1] < 0) || (nums[i+1] - nums[i] < 0 && nums[i] - nums[i-1] > 0) {

		} else {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func main() {
	//nums := []int{1,7,4,9,2,5}
	//nums := []int{1,17,5,10,13,15,10,5,16,8}
	//nums := []int{1,2,3,4,5,6,7,8,9}
	//nums := []int{0,0}
	nums := []int{0,0,0}
	fmt.Println(wiggleMaxLength(nums))
}
