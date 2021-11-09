package main

import "fmt"

func generateMatrix(n int) [][]int {
	nums := []int{}
	for i := 1; i < n * n; i++ {
		nums = append(nums, i)
	}
	matrix := make([][]int, n, n)
	tmp := []int{}
	for i := 0; i < n; i++ {
		tmp = append(tmp, 0)
	}
	for i := 0; i < n; i++ {
		matrix = append(matrix, matrix...)
	}
	fmt.Println(matrix)
	col := n
	row := 0
	for col > 0 {
		n = nums[0]
		if len(nums) == 0 {
			break
		}
		nums = nums[1:]
		for c := 0; c < col; c++ {
			matrix[row][c] = n
			col = c
		}
		for r := 0; r < row; r++ {
			matrix[r][col] = n
			row = r
		}
	}
	return matrix
}

func main() {
	n := 3
	fmt.Println(generateMatrix(n))
}
