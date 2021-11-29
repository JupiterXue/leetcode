package main

import (
	"fmt"
)

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minimumDeletions(nums []int) int {
	min := nums[0]
	minIdx := 0
	max := nums[0]
	maxIdx := 0
	for i, v := range nums {
		if min > v {
			min = v
			minIdx = i
		}
		if max < v {
			max = v
			maxIdx = i
		}
	}
	fmt.Println(minIdx)
	fmt.Println(maxIdx)
	if minIdx == maxIdx {
		return 1
	} else {
		mid := len(nums) / 2
		if (minIdx <= mid && maxIdx <= mid) ||( minIdx > mid && maxIdx > mid) {
			fmt.Println(1)
			return minIdx + maxIdx
		} else {
			fmt.Println(2)
			minIdx := getMin(minIdx, len(nums)-minIdx)
			maxIdx := getMin(maxIdx, len(nums)-maxIdx)
			return minIdx + maxIdx + 1
		}
	}
}

func main() {
	//nums := []int{2,10,7,5,4,1,8,6}
	//nums := []int{101}
	//nums := []int{0, -4, 19, 1, 8, -2, -3, 5}
	//nums := []int{48,-49,-67,18,-59,-56,47,-26,-24,-73,-96,27,-2,-45}
	nums := []int{-1,-53,93,-42,37,94,97,82,46,42,-99,56,-76,-66,-67,-13,10,66,85,-28}
	fmt.Println(minimumDeletions(nums))
}
