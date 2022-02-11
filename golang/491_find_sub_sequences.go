package main

import (
	"fmt"
)

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var path []int
	var backTracking func(startIndex int)
	backTracking = func(startIndex int) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		history := [201]int{} // 记录本
		for i := startIndex; i < len(nums); i++ {
			if len(path) > 0 && nums[i] < path[len(path)-1] || history[nums[i] + 100] == 1 {
				continue
			}
			history[nums[i] + 100] = 1 // 表示本层使用过该元素了
			path = append(path, nums[i])
			backTracking(i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return res
}

func main() {
	//nums := []int{4,6,7,7}
	//nums := []int{4,4,3,2,1}
	nums := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	fmt.Println(findSubsequences(nums))
}
