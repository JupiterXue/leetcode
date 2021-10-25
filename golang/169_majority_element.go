package main

import "fmt"

func majorityElement(nums []int) int {
	numsMap := make(map[int]int)
	numsHalfLen := len(nums) / 2
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			numsMap[v]++
		} else {
			numsMap[v] = 1
		}
	}

	//fmt.Println(numsMap)
	//fmt.Println(numsHalfLen)

	for k, v := range numsMap {
		if v > numsHalfLen {
			return k
		}
	}
	return 0
}

func main() {
	//nums := []int{3,2,3}
	//nums := []int{2,2,1,1,1,2,2}
	nums := []int{3,3,4}
	fmt.Println(majorityElement(nums))
}
