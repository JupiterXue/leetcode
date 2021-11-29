package main

import (
	"fmt"
)

func sumNums(partNums []int) int {
	count := 0
	for _, v := range partNums {
		count += v
	}
	return count
}

func getAverages(nums []int, k int) []int {
	res := []int{}
	size := k * 2 + 1
	if k > len(nums) {
		return []int{-1}
	}
	for i := 0; i < k; i++ {
		res = append(res, -1)
	}
	for i := k; i < len(nums) - k; i++ {
		partNums := nums[i-k:i+k+1]
		fmt.Println(partNums)
		sum := sumNums(partNums)
		res = append(res, sum/size)
	}
	for i := len(nums); i > len(nums) - k; i-- {
		res = append(res, -1)
	}
	return res
}

func main() {
	//nums := []int{7,4,3,9,1,8,5,2,6}
	nums := []int{10000}
	k := 0
	//k := 111
	fmt.Println(getAverages(nums, k))
}