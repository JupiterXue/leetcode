package main

import "fmt"

// 1. 自己，贪婪，AC！
//func candy(ratings []int) int {
//	res := make([]int, len(ratings))
//	for i := 0; i < len(res); i++ {
//		res[i] = 1
//	}
//	for i := 1; i < len(ratings); i++ {
//		if ratings[i] > ratings[i-1] {
//			res[i] = res[i-1] + 1
//		}
//	}
//	for i := len(ratings)-2; i >= 0; i-- {
//		if ratings[i] > ratings[i+1] {
//			if res[i] - res[i+1] <= 1 {
//				res[i] = res[i+1] + 1
//			}
//		}
//	}
//	count := 0
//	fmt.Println(res)
//	for _, v := range res {
//		count += v
//	}
//	return count
//}

// 2. 随想录，贪婪，AC！
func candy(ratings []int) int {
	need := make([]int, len(ratings))
	sum := 0
	var findMax func(need1, need2 int) int
	findMax = func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		} else {
			return num2
		}
	}
	// 1. 初始化数值
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	// 2. 从左往右，右边大于左边就+1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 3. 从右往左，左边大于右边就+1，需要判断
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] >ratings[i] {
			need[i-1] = findMax(need[i-1], need[i]+1)
		}
	}
	// 计算总糖果数
	for i := 0; i < len(ratings); i++ {
		sum += need[i]
	}
	return sum
}


func main() {
	ratings := []int{1,0,2}
	//ratings := []int{1,2,2}
	//ratings := []int{1,3,2,2,1}
	//ratings := []int{1,2,87,87,87,2,1}
	//ratings := []int{1,3,4,5,2}
	fmt.Println(candy(ratings))
}
