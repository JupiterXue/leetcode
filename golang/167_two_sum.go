package main

import "fmt"

func twoSum2(numbers []int, target int) []int {
	res := []int{}
	for i := 0; i < len(numbers) - 1 ; i++ {
		for j := i+1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == target {
				res = append(res, i+1)
				res = append(res, j+1)
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{3,22,1}
	target := 3
	fmt.Println(twoSum2(nums, target))
}
