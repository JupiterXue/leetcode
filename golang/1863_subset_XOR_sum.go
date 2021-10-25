package main

import "fmt"

var ans = 0

//func dfsXOR(nums []int, i , XOR int) {
//	if i == len(nums) {
//		ans += XOR
//		return
//	}
//
//	dfsXOR(nums, i+1, XOR)
//	dfsXOR(nums, i+1, XOR^nums[i])
//}
//
//func subsetXORSum(nums []int) int {
//	dfsXOR(nums, 0, 0)
//	return ans
//}

//var res = 0
//var depth = 0
//
//func dfsXOR(index, curSum int, nums []int) {
//	if index == depth {
//		res += curSum
//		return
//	}
//
//	dfsXOR(index + 1, curSum, nums)
//	dfsXOR(index + 1, curSum ^ nums[index], nums)
//}
//
//func subsetXORSum(nums []int) int {
//	depth = len(nums)
//	dfsXOR(0, 0, nums)
//	return res
//}

func subsetXORSum(nums []int) int {
	res := 0
	lenNnums := len(nums)
	// << 符号，将一个运算对象的各二进制位全部左移若干位，右补 0
	for i := 0; i < (1 << lenNnums); i++ {
		tmpRes := 0
		for j := 0; j < lenNnums; j++ {
			if i & (1 << j) > 0 {
				tmpRes ^= nums[j]
			}
		}
		res += tmpRes
	}
	return res
}



//func XOR(partNums []int) int {
//	res := 0
//	for _, v := range partNums {
//		res = res ^ v
//	}
//	return res
//}
//
//func subsetXORSum(nums []int) int {
//	if len(nums) == 0{
//		return 0
//	}
//
//	valueList := []int{}
//	for i := 0; i < len(nums); i++ {
//		for j := i; j < len(nums); j++ {
//			value := XOR(nums[i:j+1])
//			valueList = append(valueList, value)
//			fmt.Print(nums[i:j+1])
//		}
//		fmt.Println()
//	}
//
//	fmt.Println(valueList)
//
//	res := 0
//	for _, v := range valueList {
//		res += v
//	}
//
//	return res
//}

func main() {
	nums := []int{5, 1, 6}
	fmt.Println(1 << 3)
	fmt.Println(subsetXORSum(nums))
}
