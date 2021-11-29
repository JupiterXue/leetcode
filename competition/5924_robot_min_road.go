package main

import (
	"fmt"
	"math"
)

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	//x, y := homePos[0] -  startPos[0], homePos[1] - startPos[1]
	x, y := homePos[0], homePos[1]
	count := 0
	for i := 0; i < int(math.Abs(float64(homePos[0]-startPos[0]))); i++ {
		x++
		if x >= len(rowCosts) {
			break
		}
		count +=  rowCosts[x]
	}
	for i := 0; i < int(math.Abs(float64(homePos[1]-startPos[1]))); i++ {
		y++
		fmt.Println(y)
		fmt.Println(colCosts)
		if y >= len(colCosts) {
			break
		}
		count += colCosts[y]
	}
	return count
}

func main() {
	//startPos := []int{1,0}
	//homePos := []int{2,3}
	startPos := []int{5,5}
	homePos := []int{5,2}
	rowCosts := []int{7,1,3,3,5,3,22,10,23}
	colCosts := []int{5,5,6,2,0,16}
	fmt.Println(minCost(startPos, homePos, rowCosts, colCosts))
}
