package main

import "fmt"

func numWays(n int, relation [][]int, k int) int {
	var queue = [][]int{{0,0}}
	for n := 0; n < k; n++ {
		tmpQueue := [][]int{}
		for len(queue) > 0 {
			next := queue[0][1]
			queue = queue[1:]
			for r := 0; r < len(relation); r++ {
				if relation[r][0] == next {
					tmpQueue = append(tmpQueue, relation[r])
				}
			}
		}
		//copy(queue, tmpQueue) // 复制后到前
		//if len(tmpQueue) == 0 {
		//	return 0
		//}
		queue = append(queue, tmpQueue...)
		fmt.Println(queue)
	}
	count := 0
	for r := 0; r < len(queue); r++ {
		if queue[r][1] == n-1 {
			count++
		}
	}
	return count
}

func main() {
	//n, k := 1, 2
	//relation := [][]int{{0,2}, {2,1}}
	//n, k := 5, 3
	//relation := [][]int{{0,2},{2,1},{3,4},{2,3},{1,4},{2,0},{0,4}}
	//n, k := 3, 1
	//relation := [][]int{{0,2},{0,1},{1,2},{2,1},{2,0},{1,0}}
	n, k := 3, 5
	relation := [][]int{{0,1},{0,2},{2,1},{1,2},{1,0},{2,0}}
	fmt.Println(numWays(n, relation, k))
}
