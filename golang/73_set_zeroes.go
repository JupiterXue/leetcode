package main

import "fmt"

//func setZeroes(matrix [][]int)  {
//	rowLen := len(matrix)
//	colLen := len(matrix[0])
//	location := [][]int{}
//	for r := 0; r < rowLen; r++ {
//		for c := 0; c < colLen; c++ {
//			if matrix[r][c] == 0 {
//				location = append(location, []int{r,c})
//			}
//		}
//	}
//	fmt.Println(location)
//	for r := 0; r < len(location); r++ {
//		for c := 0; c < len(location[0]); c++ {
//			for i := 0; i < rowLen; i++ { // 行置零
//				matrix[i][location[r][c]] = 0
//			}
//			for i := 0; i < colLen; i++ {
//				matrix[location[r][c]][i] = 0 // 列置零
//			}
//		}
//
//	}
//	//fmt.Println(location)
//	//fmt.Println(matrix)
//}

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v ==0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}

func main() {
	//matrix := [][]int{{1,1,1},{1,0,1},{1,1,1}}
	matrix := [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
