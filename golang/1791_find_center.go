package main

import "fmt"

func findCenter(edges [][]int) int {
	numMap := make(map[int]int)
	for r := 0; r < len(edges); r++ {
		n := 0
		if ok := numMap[edges[r][n]]; ok > 0 {
			numMap[edges[r][n]]++
		} else {
			numMap[edges[r][n]] = 0
		}

		n = 1
		if ok := numMap[edges[r][n]]; ok > 0 {
			numMap[edges[r][n]]++
		} else {
			numMap[edges[r][n]] = 1
		}
	}

	maxK := 0
	maxV := 0
	for k, v := range numMap {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}
	return maxK
}

func main() {
	//edges := [][]int{{1,2},{2,3},{4,2}}
	edges := [][]int{{1,2},{5,1},{1,3},{1,4}}
	fmt.Println(findCenter(edges))
}
