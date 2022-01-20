package main

import "fmt"

//func generateMatrix(n int) [][]int {
//	nums := []int{}
//	for i := 1; i < n * n; i++ {
//		nums = append(nums, i)
//	}
//	matrix := make([][]int, n)
//	tmp := []int{}
//	for i := 0; i < n; i++ {
//		tmp = append(tmp, 0)
//	}
//	for i := 0; i < n; i++ {
//		matrix[i] = tmp
//	}
//	col := n
//	row := 0
//	for col > 0 {
//		n = nums[0]
//		if len(nums) == 0 {
//			break
//		}
//		nums = nums[1:]
//		for c := 0; c < col; c++ {
//			matrix[row][c] = n
//			col = c
//		}
//		for r := 0; r < row; r++ {
//			matrix[r][col] = n
//			row = r
//		}
//	}
//	return matrix
//}

//type pair struct {
//	x, y int
//}
//
//var dirs = []pair {{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上
//
//func generateMatrix(n int) [][]int {
//	matrix := make([][]int, n)
//	for i := range matrix {
//		matrix[i] = make([]int, n)
//	}
//
//	row, col, dirIdx := 0, 0, 0
//	for i:= 1; i <= n*n; i++ {
//		matrix[row][col] = i
//		dir := dirs[dirIdx]
//		if r, c := row + dir.x, col + dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
//			dirIdx = (dirIdx + 1) % 4 // 顺时针旋转至下一个方向
//			dir = dirs[dirIdx]
//		}
//		row += dir.x
//		col += dir.y
//	}
//	return matrix
//}

func generateMatrix(n int) [][]int {
	res := make([][]int, n, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	row, col := 0, 0
	loop := n / 2
	mid := n / 2
	count := 1
	size := 1
	i, j := 0, 0
	for loop > 0 {
		loop--
		i, j = row, col
		for j := col; j < col + n - size; j++ {
			res[row][j] = count
			count++
		}
		fmt.Println(res)
		j = col + n - size

		for i := row; i < row + n - size; i++ {
			res[i][j] = count
			count++
		}
		i = row + n - size

		for ; j > col; j-- {
			res[i][j] = count
			count++
		}
		j = col

		for ; i > row; i-- {
			res[i][j] = count
			count++
		}
		i = row

		row++
		col++
		size += 2
	}

	if n % 2 == 1 {
		res[mid][mid] = count
	}
	return res
}

func main() {
	n := 4
	fmt.Println(generateMatrix(n))
}
