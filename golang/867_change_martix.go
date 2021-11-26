package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i := range res {
		res[i] = make([]int, len(matrix))
		for j := range res[i] {
			res[i][j] = -1
		}
	}
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			res[c][r] = matrix[r][c]
		}
	}
	return res
}

func main() {
	//matrix := [][]int{{1,2,3}, {4,5,6},{7,8,9}}
	matrix := [][]int{{1,2,3}, {4,5,6}}
	fmt.Println(transpose(matrix))
}
