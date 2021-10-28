package main

import (
	"fmt"
	"sort"
)

func subSort(array []int) []int {
	res := []int{-1,-1}
	sortedArray := []int{}
	sortedArray = append(sortedArray, array...)
	sort.Ints(sortedArray)
	fmt.Println(array)
	fmt.Println(sortedArray)
	for i := 0; i < len(array); i++ {
		if res[0] != -1 {
			if array[i] == sortedArray[i] {
				ensure := true
				for j := i+1; j < len(array); j++ {
					if array[j] != sortedArray[j] {
						ensure = false
						break
					}
				}
				if ensure {
					res[1] = i-1
					break
				}
			}
			continue
		}
		if array[i] != sortedArray[i] {
			res[0] = i
		}
	}
	if res[0] != -1 && res[1] == -1 {
		res[1] = len(array)-1
	}
	return res
}

func main() {
	array := []int{1,2,4,7,10,11,7,12,6,7,16,18,19}
	fmt.Println(subSort(array))
}
