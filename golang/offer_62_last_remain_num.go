package main

import "fmt"

// 1. 递归。超时
//func recursionDelete(nums []int, m, deleteIdx int) []int {
//	if len(nums) == 1 {
//		return nums
//	}
//	deleteIdx = (m + deleteIdx) % len(nums)
//	nums = append(nums[:deleteIdx], nums[deleteIdx+1:]...)
//	return recursionDelete(nums, m, deleteIdx-1)
//}
//
//func lastRemaining(n int, m int) int {
//	var nums []int
//	for i := 0; i < n; i++ {
//		nums = append(nums, i)
//	}
//
//	deleteIdx := -1
//	nums = recursionDelete(nums, m, deleteIdx)
//	return nums[0]
//}

// 2. 数学方法，AC！
//func lastRemaining(n int, m int) int {
//	start := 0
//	for i := 2; i <= n; i++ {
//		start = (start + m) % i
//	}
//	return start
//}

// 3. 递归，AC！
func lastNum(n, m, p, i int) int {
	i++
	if i > n {
		return p
	}
	return lastNum(n, m, (p + m) % i, i)
}

func lastRemaining(n int, m int) int {
	p := 0
	i := 1
	p = lastNum(n, m, p, i)
	return p
}

func main() {
	n := 10
	m := 17
	fmt.Println(lastRemaining(n, m))
}
