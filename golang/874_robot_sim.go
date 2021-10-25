package main

import "fmt"

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
	dx := []int{0,1,0,-1}
	dy := []int{1,0,-1,0}
	x, y, di := 0, 0, 0

	obstaclesSet := make(map[int]int)
	for i, line := range obstacles {
		ox := line[0] + 30000
		oy := line[1] + 30000
		obstaclesSet
	}
}


func main() {
	commands := []int{}
	obstacles := [][]int{}
	fmt.Println(commands, obstacles)
}
