package main

import (
	"fmt"
	"math"
	"reflect"
)

func checkExist(partNum []int, bigNums [][]int) bool {
	count := 0
	for _, line := range bigNums {
		if reflect.DeepEqual(line, partNum) {
			count++
		}
	}
	return count >= 2
}

func sumNums(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

func isHappy(n int) bool {
	bigNums := [][]int{}
	for {
		nums := []int{}
		for math.Floor(float64(n)) != 0 {
			tmp := n % 10
			nums = append(nums, tmp*tmp)
			n /= 10
		}
		bigNums = append(bigNums, nums)
		sum := sumNums(nums)
		if sum == 1 {
			break
		} else if checkExist(nums, bigNums) { // 说明没有找到
			return false
		}
		n = sum
	}
	return true
}

func main() {
	n := 19
	fmt.Println(isHappy(n))
}
