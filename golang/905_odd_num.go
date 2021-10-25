package main

import "fmt"

// 1. 选择排序，通过
//func sortArrayByParity(nums []int) []int {
//	res := []int{}
//	oldNums := []int{}
//	for _, v := range nums {
//		if v % 2 == 0 {
//			res = append(res, v)
//		} else {
//			oldNums = append(oldNums, v)
//		}
//	}
//
//	for _, v := range oldNums {
//		res = append(res, v)
//	}
//	return res
//}

// 2. 插入排序
func sortArrayByParity(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		currentIndex := i
		for j := i - 1; j >= 0; j-- {
			if nums[currentIndex] % 2 == 0 &&  nums[j] % 2 != 0 {
				nums[currentIndex], nums[j] = nums[j], nums[currentIndex]
				currentIndex--
			}
		}
	}
	return nums
}

// 3.

func main() {
	nums := []int{3,1,2,4}
	fmt.Println(sortArrayByParity(nums))
}
