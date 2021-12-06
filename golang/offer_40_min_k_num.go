package main

import (
	"fmt"
)

// 1. 自己写，AC！
//func getLeastNumbers(arr []int, k int) []int {
//	sort.Ints(arr)
//	return arr[:k]
//}

// 2. 重写快排，AC！性能比 1 低点
//func getLeastNumbers(arr []int, k int) []int {
//	qs(arr, 0, len(arr)-1)
//	return arr[:k]
//}
//
//func qs(arr []int, left int, right int) {
//	if left < right {
//		pivot := partitionSubArr(arr, left, right)
//		qs(arr, left, pivot-1)
//		qs(arr, pivot+1, right)
//	}
//}
//
//func partitionSubArr(arr []int, left int, right int) int {
//	pivot := left
//	index := pivot + 1 // 双指针
//	for i := index; i <= right; i++ {
//		if arr[i] < arr[pivot] {
//			arr[i], arr[index] = arr[index], arr[i]
//			index++
//		}
//	}
//	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
//	return index - 1
//}

// 3. 堆排序，题解，AC
func getLeastNumbers(arr []int, k int) []int {
	for i := len(arr); i > len(arr)-k; i-- {
		heap(arr[:i])
		arr[0], arr[i-1] = arr[i-1], arr[0]
	}
	return arr[len(arr)-k:]
}

func heap(arr []int) {
	len := len(arr)
	for i := len/2 - 1; i >= 0; i-- {
		if 2*i+1 < len && arr[i] > arr[2*i+1] {
			arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
		}
		if 2*i+2 < len && arr[i] > arr[2*i+2] {
			arr[i], arr[2*i+2] = arr[2*i+2], arr[i]
		}
	}
}

func main() {
	//arr := []int{3, 2, 1}
	arr := []int{0, 1, 2, 1}
	//k := 2
	k := 1
	fmt.Println(getLeastNumbers(arr, k))
}
