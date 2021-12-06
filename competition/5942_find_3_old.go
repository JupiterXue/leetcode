package main

import (
	"fmt"
	"sort"
)

func findEvenNumbers(digits []int) []int {
	numMap := make(map[int]int)
	for _, v := range digits {
		if ok := numMap[v]; ok > 0 {
			numMap[v]++
		} else {
			numMap[v] = 1
		}
	}

	res := []int{}
	if len(digits) == numMap[1] {
		return res
	}

	onlyOdd := true
	for k := range numMap {
		if k % 2 == 0 {
			onlyOdd = false
		}
	}
	if onlyOdd {
		return res
	}

	//for i, j, k := 0, 0, 0; i < len(digits) && j < len(digits) && k < len(digits); i, j , k = i+1, j+1, k+1 {
	resMap := make(map[int]int)
	for i := 0; i < len(digits); i++ {
		if digits[i] == 0 {
			continue
		}
		for j := 0; j < len(digits); j++ {
			for k := 0; k < len(digits); k++ {
				if i == j || i == k || j == k {
					continue
				}
				if digits[k] % 2 != 0 {
					continue
				}
				tmp := digits[i]*100 + digits[j]*10 + digits[k]
				if ok := resMap[tmp]; ok > 0 {
					resMap[tmp]++
					continue
				} else {
					resMap[tmp] = 1
				}
				res = append(res, tmp)
			}
		}
	}
	sort.Ints(res)
	return res
}

func main() {
	digits := []int{2,1,3,0}
	//digits := []int{2,2,8,8,2}
	//digits := []int{3,7,5}
	//digits := []int{0,0,0}
	fmt.Println(findEvenNumbers(digits))
}