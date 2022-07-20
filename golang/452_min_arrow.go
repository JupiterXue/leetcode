package main

import (
	"fmt"
	"sort"
)

// 1. 随想录，贪婪
func findMinArrowShots(points [][]int) int {
	res := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			res++
		} else {
			if points[i-1][1] < points[i][1] {
				points[i][1] = points[i-1][1]
			} else {
				points[i][1] = points[i][1]
			}
		}
	}
	return res
}

func main() {
	points := [][]int{{10,16},{2,8},{1,6},{7,12}}
	fmt.Println(findMinArrowShots(points))
}
