package main

import (
	"fmt"
	"math"
)

//func robotSim(commands []int, obstacles [][]int) int {
//	location := [2]int{}
//	arc := 0
//	for _, n := range commands {
//		// 先判断方向
//		if n == -1 {
//			arc
//		} else if n == -2 {
//
//		}
//
//		// 判断是否有障碍
//
//		// 再判断前进多少
//
//	}
//	return location[0]*location[0] + location[1]*location[1]
//}

func robotSim(commands []int, obstacles [][]int) int {
	// 方向：上右下左0123
	// 结果，当前方向，障碍物哈希
	result, curDir, m0bstacles := float64(0), 0, make(map[string]bool)
	// 当前位置，xy 轴各个方向移动的大小
	x, y, stepX, stepY := 0, 0, []int{0,1,0,-1}, []int{1,0,-1,0}
	for _, v := range obstacles {
		m0bstacles[fmt.Sprintf("%d-%d", v[0], v[1])] = true
	}
	for _, v := range commands {
		switch v {
		case -1:
			curDir = (curDir+1) % 4
		case -2:
			curDir = (curDir+3) % 4
		default:
			for i := 0; i < v; i++ {
				tempX, tempY := x + stepX[curDir], y + stepY[curDir]
				// 碰到障碍物，不移动
				if m0bstacles[fmt.Sprintf("%d-%d", tempX, tempY)] {
					break
				}
				x, y = tempX, tempY
				result = math.Max(float64(x*x+y*y), result)
			}
		}
	}
	return int(result)
}


func main() {
	commands := []int{}
	obstacles := [][]int{}
	fmt.Println(commands, obstacles)
}
