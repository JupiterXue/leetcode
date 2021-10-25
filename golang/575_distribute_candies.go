package main

import "fmt"

func isContain(candyTypeSet []int, v int) bool {
	for _, vs := range candyTypeSet {
		if v == vs {
			return true
		}
	}
	return false
}

func distributeCandies(candyType []int) int {
	halfLen := len(candyType) / 2
	var candyTypeSet []int
	for _, v := range candyType {
		if isContain(candyTypeSet, v) {
			continue
		} else {
			candyTypeSet = append(candyTypeSet, v)
		}
	}
	if len(candyTypeSet) > halfLen {
		return halfLen
	} else {
		return len(candyTypeSet)
	}
}

func main() {
	//candyType := []int{1,1,2,2,3,3}
	//candyType := []int{1,1,2,3}
	candyType := []int{1,1,1,1,2,2,2,3,3,3}
	fmt.Println(distributeCandies(candyType))
}
