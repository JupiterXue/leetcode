package main

import "fmt"

func containsDuplicate(nums []int) bool {
	elemMap := make(map[int]int)
	for i := range nums {
		elem := nums[i]
		if ok := elemMap[elem] ; ok > 0 {
			return true
		} else {
			elemMap[elem] = 1
		}
	}
	return false
}

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(containsDuplicate(nums))
}
