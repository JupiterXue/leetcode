package main

import "fmt"

// 1. 自己写，未通过。。
//func validPath(n int, edges [][]int, start int, end int) bool {
//	if n == 1 {
//		return true
//	}
//	quque := [][]int{}
//	for r := 0; r < len(edges); r++ {
//		if (edges[r][0] == start && edges[r][1] == end) || (edges[r][1] == start && edges[r][0] == end) {
//			return true
//		}
//		if edges[r][0] == start || edges[r][1] == start {
//			quque = append(quque, edges[r]) // 入队
//		}
//	}
//	fmt.Println(quque)
//
//	for i := 0; i < n; i++ {
//		tmpQueue := [][]int{}
//		if len(quque) > 0 {
//			next0 := quque[0][0]
//			next1 := quque[0][1]
//			quque = quque[1:] // 出队
//			for r := 0; r < len(edges); r++ {
//				if next0 == edges[r][0] {
//					if edges[r][1] == end {
//						return true
//					} else {
//						tmpQueue = append(tmpQueue, edges[r])
//					}
//				}
//				if next0 == edges[r][1] {
//					if edges[r][0] == end {
//						return true
//					} else {
//						tmpQueue = append(tmpQueue, edges[r])
//					}
//				}
//			}
//			for r := 0; r < len(edges); r++ {
//				if next1 == edges[r][0] {
//					if edges[r][1] == end {
//						return true
//					} else {
//						tmpQueue = append(tmpQueue, edges[r])
//					}
//				}
//				if next1 == edges[r][1] {
//					if edges[r][0] == end {
//						return true
//					} else {
//						tmpQueue = append(tmpQueue, edges[r])
//					}
//				}
//			}
//			quque = append(quque, tmpQueue...)
//			fmt.Println(quque)
//		}
//	}
//	return false
//}

// 2. 题解
func validPath(n int, edges[][]int, start, end int) bool {
	if start == end {
		return true
	}
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parents[x] == x {
			return x
		}
		return find(parents[x])
	}
	for _, e := range edges {
		u := find(e[0])
		v := find(e[1])
		if u != v {
			parents[u] = v
		}
	}
	return find(start) == find(end)
}

func main() {
	//n, start, end := 3, 0, 2
	//edges := [][]int{{0,1},{1,2},{2,0}}
	//n, start, end := 5, 0, 4
	//edges := [][]int{{0,4}}
	//n, start, end := 1, 0, 0
	//edges := [][]int{}
	n, start, end := 10, 7, 5
	edges := [][]int{{0,7}, {0,8},{6,1},{2,0},{0,4},{5,8},{4,7},{1,3},{3,5},{6,5}}
	fmt.Println(validPath(n, edges, start, end))
}
