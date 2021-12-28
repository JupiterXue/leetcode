package main

import "fmt"

func minMoves2(nums []int) int {
	fmt.Println(nums)
	qqs(nums, 0, len(nums)-1)
	fmt.Println(nums)
	mid := nums[len(nums)/2]
	count := 0
	for _, v := range nums {
		if mid >= v {
			count += mid - v
		} else {
			count += v - mid
		}
	}
	return count
}

func qqs(nums []int, low int, high int) {
	if low < high {
		pivotIdx := ppartition(nums, low, high)
		qqs(nums, low, pivotIdx-1)
		qqs(nums, pivotIdx+1, high)
	}
}

func ppartition(nums []int, low int, high int) int {
	pivot := nums[low]
	for low < high {
		for low < high && pivot <= nums[high] {
			high--
		}
		nums[low] = nums[high]
		for low < high && pivot >= nums[low] {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}

func main() {
	//nums := []int{1,2,3}
	nums := []int{2,3,1}
	fmt.Println(minMoves2(nums))
}
