package main

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	if len(nums1) > len(nums2) {
		for _, v := range nums2 {
			for _, v2 := range nums1 {
				if v == v2 {
					res = append(res, v)
					break
				}

			}
		}
	} else {
		for _, v := range nums1 {
			for _, v2 := range nums2 {
				if v == v2 {
					res = append(res, v)
					break
				}
			}
		}
	}
	sort.Ints(res)
	for i := 0; i < len(res) - 1; i++ {
		if res[i] == res[i+1] {
			res = append(res[:i], res[i+1:]...)
			i--
		}
	}
	return res
}

func main() {
	nums := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersection(nums, nums2))
}
