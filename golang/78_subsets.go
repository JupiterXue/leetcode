package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	var backTracking func(startIndex int)
	backTracking = func(startIndex int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backTracking(i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return res
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(subsets(nums))
}
