package main

import (
	"fmt"
)

// AC!!!
//func singleNumber(nums []int) int {
//	sort.Ints(nums)
//	for i := 0; i < len(nums) - 2; {
//		if nums[i] != nums[i+1] {
//			return nums[i]
//		}
//		i++
//		i++
//	}
//	return nums[len(nums)-1]
//}

// 异或
func singleNumber(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}


func main() {
	nums := []int{2,2,1}
	//nums := []int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))
}
