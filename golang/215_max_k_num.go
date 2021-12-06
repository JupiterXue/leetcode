package main

import (
	"fmt"
)

// 1. 自己写，AC！
//func findKthLargest(nums []int, k int) int {
//	sort.Ints(nums)
//	return nums[len(nums)-k]
//}

//func findKthLargest(nums []int, k int) int {
//	QuickSort(nums, 0, len(nums)-1)
//	return nums[len(nums)-k]
//}
//
//func QuickSort(nums []int, low int, high int) {
//	if high > low {
//		// 位置划分
//		pivot := partition(nums, low, high)
//		// 左边部分排序
//		QuickSort(nums, low, pivot-1)
//		// 右边排序
//		QuickSort(nums, pivot+1, high)
//	}
//}
//
//func partition(nums []int, low int, high int) int {
//	pivot := nums[low] // 导致 low位置 为空
//	for low < high {
//		// high 指针值 >= pivot  high指针 左移
//		for low < high && pivot <= nums[high] {
//			high--
//		}
//		// 填补 low 位置空值
//		// high指针值 < pivot  high值 移动到 low 位置
//		// high位置 置空
//		nums[low] = nums[high]
//		// low指针值 <= pivot  low指针 右移
//		for low < high && pivot >= nums[low] {
//			low++
//		}
//		// 填补 high 位置空值
//		// low指针值 > pivot   low值 移动到 high 位置
//		// low位置 置空
//		nums[high] = nums[low]
//	}
//	// pivot 填补 low位置 的空值
//	nums[low] = pivot
//	return low
//}

func findKthLargest(nums []int, k int) int {
	QuickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func QuickSort(nums []int, low int, high int) {
	if low < high {
		pivot := Partition(nums, low, high)
		QuickSort(nums, low, pivot-1)
		QuickSort(nums, pivot+1, high)
	}
}

func Partition(nums []int, low int, high int) int {
	pivot := nums[low]
	for low < high {
		for low < high && pivot <= nums[high] {
			high--
		}
		nums[high] = nums[low]
		for low < high && pivot >= nums[low] {
			low++
		}
		nums[low] = nums[high]
	}
	nums[low] = pivot
	return low
}


func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	//nums := []int{3,2,3,1,2,4,5,5,6}
	k := 2
	//k := 4
	fmt.Println(findKthLargest(nums, k))
}
