package main

import (
	"fmt"
)

//func removeElement(nums []int, val int) int {
//	// 默认情况下 Golang 的数组传递是值传递.
//	// 想要对传入的数组进行修改时，可以使用指针来对数组进行操作
//	// Golang 中的切片与 Java 中的 ArrayList 集合类似，进行的是引用传递
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == val {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//		}
//	}
//	return len(nums)
//}


// 1. 暴力法， AC
//func removeElement(nums []int, val int) int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == val {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//		}
//	}
//	return len(nums)
//}

// 2. 双指针法
func removeElement(nums []int, val int) int {
	// left 作为剩余元素计数器
	left := 0
	length := len(nums)
	for right := 0; right < length; right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}


func main() {
	//make构造一个切片
	nums := []int{3,2,2,3}
	val := 2
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
	fmt.Println(len(nums))
}
