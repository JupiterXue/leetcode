package main

import "fmt"

func search(nums []int, target int) int {
	// 数组中的值互不相同
	flag := false
	index := -1
	for i, v := range nums {
		if v == target {
			flag = true
			index = i
		}
	}
	if flag == false {
		return index
	}

	return index
}

func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	fmt.Println(search(nums, target))
}
