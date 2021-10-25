package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	// 合并后数组不应由函数返回，而是存储在数组 nums1 中
	for i, v := range nums2 {
		nums1[m+i] = v
	}
	sort.Ints(nums1)
	fmt.Println(nums1)
}

func main() {
	nums := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums, m, nums2, n)
	//fmt.Println()
}
