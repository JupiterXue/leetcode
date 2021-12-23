package main

import (
	"fmt"
	"math"
)


// 自己写，fail，不知道怎么计算最远的时间
//func networkDelayTime(times [][]int, n int, k int) int {
//	count := 0
//	queue := [][]int{}
//	for r := 0; r < len(times); r++ {
//		//fmt.Println(times[r][0])
//		if times[r][0] == k {
//			queue = append(queue, times[r])
//		}
//	}
//	nums := []int{k}
//	for len(queue) > 0 {
//		minValue := 101
//		minIdx := 0
//		tmpQueue := make([][]int, len(queue))
//		copy(tmpQueue, queue)
//		queue = [][]int{} // 清空
//		for len(tmpQueue) > 0 {
//			for r := 0; r < len(tmpQueue); r++ {  // 1. 找出权重最小的节点
//				value := tmpQueue[r][2]
//				if value < minValue {
//					minValue = value
//					minIdx = r
//				}
//			}
//			next := tmpQueue[minIdx][1]
//			value := tmpQueue[minIdx][2]
//			tmpQueue = append(tmpQueue[:minIdx], tmpQueue[minIdx+1:]...)
//			if !exist(nums, next) {
//				for r := 0; r < len(times); r++ {  // 2. 更新邻居节点
//					if times[r][0] == next {
//						queue = append(queue, times[r])
//					}
//				}
//				nums = append(nums, next)  // 记录已标记节点
//				count += value  // 3. 更新邻居节点开销
//			}
//		}
//	}
//	fmt.Println(nums)
//	// 5. 计算最短路径
//	if len(nums) == n {
//		return count
//	} else {
//		return -1
//	}
//}
//
//func exist(nums []int, z int) bool {
//	for _, v := range nums {
//		if v == z {
//			return true
//		}
//	}
//	return false
//}

// 2. 题解，AC！Dijkstra
func networkDelayTime(times [][]int, n int, k int) int {
	// 需要求出节点 kk 到其余所有点的最短路，其中的最大值就是答案。
	const inf = math.MaxInt64 / 2
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		g[x][y] = t[2]
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}

	dist[k-1] = 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		for y, u := range used {
			if !u && (x == -1 || dist[y] < dist[x]) {
				x = y
			}
		}
		used[x] = true
		for y, time := range g[x] {
			dist[y] = min(dist[y], dist[x]+time)
		}
	}

	ans := 0
	for _, d := range dist {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}


func main() {
	times := [][]int{{2,1,1},{2,3,1},{3,4,1}}
	n, k := 4, 2
	//times := [][]int{{1,2,1}}
	//n, k := 2, 1
	//times := [][]int{{1,2,1}}
	//n, k := 2, 2
	//times := [][]int{{1,2,1}, {2,1,3}}
	//n, k := 2, 2
	fmt.Println(networkDelayTime(times, n, k))
}
