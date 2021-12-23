package main

import "fmt"

// 1. 自己写，没写出，动态
//func maxProduct(nums []int) int {
//	max := nums[0]
//	for i := 2; i < len(nums); i++ {
//		if nums[i] * nums[i-1] > nums[i-1] {  // 如果乘积更大，累积
//			nums[i] *= nums[i-1]
//		}
//		if nums[i] * nums[i-2] > nums[i-1] {  // 如果乘积更大，累积
//			nums[i] *= nums[i-1]
//		}
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	return max
//}

func maxProduct(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	res, max, min := nums[0], nums[0], nums[0]
	for i := 1; i < size; i++ {
		mx := max
		max = getMax(max*nums[i], getMax(nums[i], nums[i]*min))
		min = getMin(min*nums[i], getMin(nums[i], mx*nums[i]))
		res = getMax(max, res)
	}
	return res
}

func getMax(i int, i2 int) int {
	if i > i2 {
		return i
	} else {
		return i2
	}
}

func getMin(i int, i2 int) int {
	if i > i2 {
		return i2
	} else {
		return i
	}
}



func main() {
	//nums := []int{2,3,-2,4}
	nums := []int{-2,3,-4}
	fmt.Println(maxProduct(nums))
}
