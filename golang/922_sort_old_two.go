package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	res := []int{}
	oddNums := []int{}
	evenNums := []int{}
	for _, v := range nums {
		if v % 2 == 0 {
			evenNums = append(evenNums, v)
		} else {
			oddNums = append(oddNums, v)
		}
	}

	for _, v := range evenNums {
		res = append(res, v)
		res = append(res, 0)
	}

	oddIndex := 1
	for _, v := range oddNums {
		res[oddIndex] = v
		oddIndex += 2
	}

	return res
}

func main() {
	nums := []int{4,2,5,7}
	fmt.Println(sortArrayByParityII(nums))
}
