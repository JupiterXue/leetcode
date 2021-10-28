package main

import (
	"fmt"
	"math"
)

// 未通过！--r 问题
func findSquare(matrix [][]int) []int {
	// 参考：https://leetcode-cn.com/problems/max-black-square-lcci/solution/a-fei-suan-fa-mian-shi-ti-1723-zui-da-he-lap0/
	// 返回一个数组 [r, c, size]
	// 其中 r, c 分别代表子方阵左上角的行号和列号，size 是子方阵的边长。
	// 返回 c 最小的子方阵。若无满足条件的子方阵，返回空数组。
	rowLen := len(matrix)
	colLen := len(matrix[0])
	// 记录当前点 [r, c] 到右边，能探到的黑色点的距离
	maxRow := make([][]int, rowLen, colLen)
	// 记录当前点 [r, c] 到下边，能探到的黑色点的距离
	maxCol := make([][]int, rowLen, colLen)
	for r := rowLen-1; r >= 0; {
		r--
		for c := colLen-1; c >= 0;  {
			c--
			if matrix[r][c] == 0 { // 本身是黑色点，开始计算，只有本身一个黑色点的话，编程是 0
				maxRow[r][c] = 1
				if c < colLen - 1 {
					maxRow[r][c] += maxRow[r][c+1]
				}
				maxCol[r][c] = 1
				if r < rowLen - 1 {
					maxCol[r][c] += maxCol[r+1][c]
				}

			}
		}
	}

	res := make([]int, 3)
	//res := [3]int{}
	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			if matrix[r][c] == 0{
				// 到下面和右边所能探到的最大距离
				upper := int(math.Min(float64(maxRow[r][c]), float64(maxCol[r][c])))
				bound := res[2]
				for size := upper; size > bound; size-- {
					if maxRow[r+size-1][c] >= size && maxCol[r][c+size-1] >= size {
						res[0] = r
						res[1] = c
						res[2] = size
						break
					}
				}
			}
		}
	}

	if res[2] == 0 {
		return []int{}
	} else {
		return res
	}
}

func main() {
	matrix := [][]int{{0,1,1},{1,0,1},{1,1,0}}
	fmt.Println(findSquare(matrix))
}
