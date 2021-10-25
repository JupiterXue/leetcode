package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	for n := 1; n <= k; n++ {
		res := [][]int{}
		last := grid[len(grid)-1][len(grid[0])-1]
		for i := 0; i < len(grid); i++ {
			tmp := []int{last}
			tmp = append(tmp, grid[i]...)
			last = grid[i][len(grid[i])-1]
			res = append(res, tmp[:len(tmp)-1])
		}
		//fmt.Println(last)
		grid = res
	}
	return grid
}

func main() {
	gird := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	k := 2
	fmt.Println(shiftGrid(gird, k))
}
