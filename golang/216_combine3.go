package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int
	var backTracking func(n, k, startIndex int)
	var sumPath func(path []int) int
	backTracking = func(n, k, startIndex int) {
		if len(path) == k {
			if sumPath(path) == n {
				tmp := make([]int, k)
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := startIndex; i <= 9; i++ {
			path = append(path, i)
			backTracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	sumPath = func(path []int) int {
		sum := 0
		for _, v := range path {
			sum += v
		}
		return sum
	}
	backTracking(n, k, 1)
	return res
}

func main() {
	n, k := 9, 45
	res := combinationSum3(n, k)
	fmt.Println(res)
}