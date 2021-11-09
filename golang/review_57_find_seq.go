package main

import (
	"fmt"
	"math"
)

//func genNums(start, end int) []int {
//	tmp := []int{}
//	for i := start; i <= end; i++ {
//		tmp = append(tmp, i)
//	}
//	return tmp
//}
//
//func sumNums(nums []int) int {
//	total := 0
//	for _, v := range nums {
//		total += v
//	}
//	return total
//}
//
//func findContinuousSequence(target int) [][]int {
//	res := [][]int{}
//	for i := 1; i < target / 2 + 1; i++ {
//		for j := i; j < target; j++ {
//			nums := genNums(i, j)
//			if sumNums(nums) == target {
//				res = append(res, nums)
//			}
//		}
//	}
//	return res
//}

func genNums(start, end int) []int {
	tmp := []int{}
	for i := start; i < end; i++ {
		tmp = append(tmp, i)
	}
	return tmp
}

func findContinuousSequence(target int) [][]int {
	i, j, sum := 1, 1, 0
	res := [][]int{}

	for i < int(math.Floor(float64(target))) {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			arr := genNums(i, j)
			res = append(res, arr)
			sum -= i
			i++
		}
	}
	return res
}


func main() {
	target := 9
	fmt.Println(findContinuousSequence(target))
}
