package main

import (
	"fmt"
	"reflect"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var path []int
	var isExists func(path []int) bool
	isExists = func(path []int) bool {
		for _, linePath := range res {
			if reflect.DeepEqual(linePath, path) {
				return true
			}
		}
		return false
	}
	var backTracking func(n, k, startIndex int)
	backTracking = func(n, k, startIndex int) {
		if !isExists(path) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backTracking(0,0,i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0,0,0)
	return res
}

func main() {
	//nums := []int{1,2,2}
	nums := []int{0}
	fmt.Println(subsetsWithDup(nums))
}
