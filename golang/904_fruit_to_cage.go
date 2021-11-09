package main

import "fmt"

//// 根据题意写判断函数
//func check(n Type)  {/*statement*/}
//// 形参列表根据题意变换
//func slidingWindow(nums []int) {
//	n := len(nums)
//	// 使用 i 指针遍历整个数组
//	for i, j := 0, 0; i < n; i++ {
//		// 调整 j 指针使得 [i, j] 符合题意
//		for j <= i && check() {
//			/* statement */
//			j++
//		}
//	}
//}

func totalFruit(fruits []int) int {
	// 建立哈希表，记录已在篮子中水果个数（区间内元素种类数）
	hashMap := map[int]int{}
	i, j ,res, n := 0, 0, 0, len(fruits)
	// i 指针遍历数组
	for ; i < n; i++ {
		// 为得到以 i 为结尾的区间，i 必须取到
		hashMap[fruits[i]]++
		// 当哈希表长度为 3 ，区间元素种类有 3 个，不符
		for j <= i && len(hashMap) == 3 {
			// j 指针右移，将 j 只想水果数 - 1
			hashMap[fruits[j]]--
			// 表示篮子中没有这种水果了，删除
			if hashMap[fruits[j]] == 0 {
				delete(hashMap, fruits[j])
			}
			j++
		}
		// 每次循环后，得到以 i 结尾的，符合题意的最长区间
		// 再通过筛选每次的结果，得到最优解
		if res < i - j + 1 {
			res = i - j + 1
		}
	}
	return res
}

func main() {
	fmt.Println()
}
