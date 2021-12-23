package main

import "fmt"

// 1. 自己写，超时
//func findDuplicates(nums []int) []int {
//	res := []int{}
//	for len(nums) > 0 {
//		count := 1
//		num := nums[0]
//		nums = nums[1:]
//		for i := 0; i < len(nums); i++ {
//			if num == nums[i] {
//				count++
//				nums = append(nums[:i], nums[i+1:]...)
//			}
//		}
//		if count == 2 {
//			res = append(res, num)
//		}
//	}
//	return res
//}

// 2. 题解，布隆过滤器
func findDuplicates(nums []int) []int {
	res := []int{}
	nums2 := make([]int,len(nums)+1)
	fmt.Println(nums2)
	for _, v := range nums {
		if nums2[v] == 1 {
			res = append(res, v)
			continue
		}
		nums2[v] = 1
	}
	return res
}


func main() {
	nums := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDuplicates(nums))
}
