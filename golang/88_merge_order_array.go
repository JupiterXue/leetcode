package main

import "fmt"

// 1. 偷鸡做法，AC
//func merge(nums1 []int, m int, nums2 []int, n int)  {
//	// 合并后数组不应由函数返回，而是存储在数组 nums1 中
//	for i, v := range nums2 {
//		nums1[m+i] = v
//	}
//	sort.Ints(nums1)
//	fmt.Println(nums1)
//}

func searchMinIndex(nums []int) int {
	Idx := 0
	value := nums[0]
	for i, v := range nums {
		if v < value {
			Idx = i
			value = v
		}
	}
	return Idx
}

// 2. 选择排序
func merge(nums1 []int, m int, nums2 []int, n int)  {
	//bigNums := append(nums1[:m], nums2...) // 天啊，这是浅拷贝，搞死我了
	bigNums := make([]int, m+n)
	copy(nums1[m:], nums2)  // 深拷贝
	copy(bigNums, nums1)    // 深拷贝
	i := 0
	for len(bigNums) > 0 {
		minIdx := searchMinIndex(bigNums)
		minValue := bigNums[minIdx]
		bigNums = append(bigNums[:minIdx], bigNums[minIdx+1:]...)
		nums1[i] = minValue
		i++
	}
	fmt.Println(nums1)
}

func main() {
	nums := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums, m, nums2, n)
	fmt.Println(nums)
}
