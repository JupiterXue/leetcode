package main

import (
	"fmt"
	"sort"
)

func merge2(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] < intervals[i][0] {
			res = append(res, intervals[0])
			intervals = append(intervals[:i-1], intervals[i:]...)
			i--
		} else {
			if intervals[i-1][1] < intervals[i][1] {
				intervals[i-1][1] = intervals[i][1]
			}
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
		}
	}
	res = append(res, intervals[len(intervals)-1])
	return res
}

func main() {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	//intervals := [][]int{{1,4},{4,5}}
	//intervals := [][]int{{1,4},{2,3}}
	fmt.Println(merge2(intervals))
}
