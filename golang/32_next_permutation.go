package main

import "fmt"

func reverse2(numsPart []int) {
	for i, n := 0, len(numsPart); i < n/2; i++ {
		numsPart[i], numsPart[n-1-i] = numsPart[n-1-i], numsPart[i]
	}
}

func nextPermutation(nums []int)  {
	numsLen := len(nums)
	i := numsLen - 2
	// 逆序比较，直到找到前一个小于后一个的
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := numsLen - 1
		// 正序比较，直到找到前一个小于后一个的
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 数组反转
	reverse2(nums[i+1:])
}

func main() {
	nums := []int{1,2,3}
	nextPermutation(nums)
	fmt.Println(nums)
}
