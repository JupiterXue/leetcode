package main

import "fmt"

func binarySearch(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		midValue := nums[mid]
		if midValue == target {
			return true
		} else if midValue < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	for row := 0; row < len(matrix); row++ {
		for binarySearch(matrix[row], target) {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1,3,5,7}, {10,11,16,20}, {23,30,34,60}}
	target := 13
	fmt.Println(searchMatrix(matrix, target))
}
