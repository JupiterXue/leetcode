package main

import (
	"fmt"
	"strconv"
)

// 1. 自己写，暴力，超时
//func monotoneIncreasingDigits(n int) int {
//	for i := n; i > 0 ; i-- {
//		nums := []int{}
//		tmp := i
//		for tmp > 0 {
//			nums = append(nums, tmp%10)
//			tmp /= 10
//		}
//		if upCheck(nums) {
//			return i
//		}
//		//fmt.Println(i)
//		//fmt.Println(nums)
//	}
//	return 1
//}
//
//func upCheck(nums []int) bool {
//	for i := 1; i < len(nums); i++ {
//		if nums[i] > nums[i-1] {
//			return false
//		}
//	}
//	return true
//}

// 另一个作死
//func monotoneIncreasingDigits(n int) int {
//	var checkIncreasing func(tmp int) bool
//	checkIncreasing = func(num int) bool {
//		numS := strconv.Itoa(num)
//		if len(numS) == 1 {
//			return true
//		}
//		pre := numS[0]
//		for i := 1; i < len(numS); i++ {
//			if numS[i] >= pre {
//				pre = numS[i]
//			} else {
//				return false
//			}
//		}
//		return true
//	}
//	for i := n; i >= 0; i-- {
//		if checkIncreasing(i) {
//			return i
//		}
//	}
//	return -1
//}

// 1. 自己写，看了题解思路
//func monotoneIncreasingDigits(n int) int {
//	tmp := n
//	nums := []int{}
//	for tmp > 0 {
//		nums = append(nums, tmp%10)
//		tmp /= 10
//	}
//	for i, j := 0, len(nums)-1; i < len(nums)/2; i, j = i+1, j-1 {
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//	for {
//		isUp := true
//		for i := 1; i < len(nums); i++ {
//			if nums[i] < nums[i-1] {
//				nums[i-1]--
//				isUp = false
//				for j := i; j < len(nums); j++ {
//					nums[j] = 9
//				}
//			}
//		}
//		if isUp {
//			break
//		}
//	}
//	res := 0
//	pow := len(nums)-1
//	for _, v := range nums {
//		res += v * int(math.Pow(10, float64(pow)))
//		pow--
//	}
//	return res
//}

// 2. 随想录，贪婪，AC
func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}
	arr := []byte(strconv.Itoa(n))
	length := len(arr)
	for i := length-1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			arr[i-1] -= 1
			for j := i; j < length; j++ {
				arr[j] = '9'
			}
		}
	}
	fmt.Println(string(arr))
	res, _ := strconv.Atoi(string(arr))
	return res
}


func main() {
	n := 10
	//n := 20
	//n := 1234
	//n := 332 // 299
	//n := 12332 // 12299
	//n := 963856657
	fmt.Println(monotoneIncreasingDigits(n))
}
