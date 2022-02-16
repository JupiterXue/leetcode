package main

import "fmt"

// 1. 自己写，逻辑没想清楚
//func jump(nums []int) int {
//	count := 1
//	cur := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] > cur - 1 {
//			cur = nums[i]
//			count++
//		}
//	}
//	return count
//}

// 2. 随想录，贪婪
func jump(nums []int) int {
	res := 0
	if len(nums) <= 1 {
		return res
	}
	cur, next := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] + i > next {
			next = nums[i] + i
		}
		if i == cur {
			if cur != len(nums) - 1 {
				res++
				cur = next
				if next >= len(nums) - 1 {
					break
				}
			} else {
				break
			}
		}
	}
	return res
}

func main() {
	//nums := []int{2,3,1,1,4}
	nums := []int{2,3,0,1,4}
	//nums := []int{3,2,1}
	fmt.Println(jump(nums))
}
