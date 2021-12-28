package main

import "fmt"

//func partition(arr []int,left, right int) int {
//	pivot := left
//	index := pivot + 1
//	for i := index; i <= right; i++ {
//		if arr[i] < arr[pivot] {
//			arr[i], arr[index] = arr[index], arr[i]
//			index++
//		}
//	}
//	arr[pivot], arr[index - 1] = arr[index - 1], arr[pivot]
//	return index - 1
//
//}
//
//func qs(arr []int, left, right int) []int {
//	if left < right {
//		partitionIndex := partition(arr, left, right)
//		qs(arr, left, partitionIndex - 1)
//		qs(arr, partitionIndex + 1, right)
//	}
//	return arr
//}
//
//func sortArray(nums []int) []int {
//	return qs(nums, 0, len(nums)-1)
//}

//func partion(nums []int, left, right int) int {
//	for left < right {
//		for left < right && nums[left] <= nums[right] {
//			right--
//		}
//
//		if left < right {
//			nums[left], nums[right] = nums[right], nums[left]
//			left++
//		}
//
//		for left < right && nums[left] <= nums[right] {
//			left++
//		}
//
//		if left < right {
//			nums[left], nums[right] = nums[right], nums[left]
//			right--
//		}
//	}
//	return left
//}
//
//func quickSort(nums []int, left, right int)  {
//	if left < right {
//		pivot := partion(nums, left, right)
//		quickSort(nums, left, pivot - 1)
//		quickSort(nums, pivot + 1, right)
//	}
//}

//	1. 冒泡，超时了。。。
//func sortArray(nums []int) []int {
//	// 升序排列
//	//for i := 0; i < len(nums) - 1; i++ {
//	//	for j := 0; j < len(nums) - i -1; j++ {
//	//		if nums[j] > nums[j+1] {
//	//			nums[j], nums[j+1] = nums[j+1], nums[j]
//	//		}
//	//	}
//	//}
//	//return nums
//
//
//}

// 2. 选择排序，超时
//func sortArray(nums []int) []int {
//
//	res := []int{}
//	for len(nums) != 0 {
//		minIndex := 0
//		min := nums[0]
//		for i, v := range nums {
//			if v < min {
//				min = v
//				minIndex = i
//			}
//		}
//		res = append(res, nums[minIndex])
//		nums = append(nums[:minIndex], nums[minIndex+1:]...)
//	}
//	return res
//}

// 3. 插入排序，超时
//func sortArray(nums []int) []int {
//	for i := 1; i < len(nums); i++ {
//		curentIndex := i
//		for j := i - 1; j >= 0; j-- {
//			if nums[curentIndex] < nums[j] { // 后面的比前面小要交换位置
//				nums[curentIndex], nums[j] = nums[j], nums[curentIndex]
//				curentIndex--
//			}
//		}
//	}
//	return nums
//}

// 4. 快排，AC！
//func sortArray(nums []int) []int {
//	quickSort(nums, 0, len(nums)-1)
//	return nums
//}
//
//func quickSort(nums []int, left int, right int) {
//	if left < right {
//		pivot := partition(nums, left, right)
//		quickSort(nums, left, pivot - 1)
//		quickSort(nums, pivot + 1, right)
//	}
//}
//
//func partition(nums []int, left int, right int) int {
//	pivot := left
//	index := pivot + 1
//	for i := index; i <= right; i++ {
//		if nums[i] < nums[pivot] {
//			nums[i], nums[index] = nums[index], nums[i]
//			index++
//		}
//	}
//	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
//	return index - 1
//}

func sortArray(nums []int) []int {
	qs(nums, 0, len(nums)-1)
	return nums
}

func qs(nums []int, low int, high int) {
	if low < high {
		pivotIdx := partition(nums, low, high)
		qs(nums, low, pivotIdx-1)
		qs(nums, pivotIdx+1, high)
	}
}

func partition(nums []int, low int, high int) int {
	pivot := nums[low]
	for low < high {
		for low < high && pivot <= nums[high] {
			high--
		}
		nums[low] = nums[high]
		for low < high && pivot >= nums[high] {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}




//func partition(arr []int,left, right int) int {
//	pivot := left
//	index := pivot + 1
//	for i := index; i <= right; i++ {
//		if arr[i] < arr[pivot] {
//			arr[i], arr[index] = arr[index], arr[i]
//			index++
//		}
//	}
//	arr[pivot], arr[index - 1] = arr[index - 1], arr[pivot]
//	return index - 1
//}

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}
