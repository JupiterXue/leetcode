package main

//void backtracking（参数）{
//if (终止条件)
//	收集结果
//	return
//for 集合的元素集
//	处理节点
//	递归函数（添加操作）
//	回溯函数（撤销操作）
//}


//// 还可以作为内置函数
//func backTracking(n, k, index int, nums[]int) {
//	if len(nums) == k {
//		tmp := make([]int, k)
//		copy(tmp, nums) // 深拷贝
//		res = append(res, tmp)
//		return
//	}
//
//	for i := index; i <= n; i++ {
//		backTracking(n, k, i+1, append(nums, i))
//	}
//}

// 还可以作为内置函数

//var res [][]int
//func backTracking(begin, end, k int) [][]int {
//	var res [][]int
//	for i := begin; i <= end; i++ {
//		if 1 == k {
//			res = append(res, []int{i})
//			continue
//		}
//		sub := backTracking(i+1, end, k-1)
//		for _, j := range sub {
//			res = append(res, append([]int{i}, j...))
//		}
//	}
//	return res
//}
//
//func combine(n int, k int) [][]int {
//	return backTracking(1, n, k)
//}
//
//func main() {
//	n := 4
//	k := 2
//	fmt.Println(combine(n, k))
//}


func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var backTracking func(n, k, startIndex int)
	backTracking = func(n, k, startIndex int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := startIndex; i <= n - (k-len(path)) + 1; i++ {
			path = append(path, i)
			backTracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(n, k, 1)
	return res
}