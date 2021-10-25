package main

import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	//arrSet := []int{}
	numList := []int{}

	//for _, v := range arr {
	//	if _, ok := arrSet[v]; ok {
	//
	//	}
	//}

	for i := 0; i < len(arr); i++ {
		count := 0
		for _, v := range arr {
			if arr[i] == v {
				count += 1
			}
		}
		numList = append(numList, count)
	}

	sort.Ints(numList)
	for i := 0; i < len(numList) - 1; i++ {
		if numList[i] == numList[i+1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1,2,2,1,1,3}
	fmt.Println(uniqueOccurrences(arr))
}
