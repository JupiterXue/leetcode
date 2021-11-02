package main

import (
	"fmt"
)

func max_num(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	i, j, res := 0, len(height)-1, 0
	for i < j {
		if height[i] < height[j] {
			res = max_num(res, height[i] * (j-i))
			i++
		} else {
			res = max_num(res, height[j] * (j-i))
			j--
		}
	}
	return res
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}
