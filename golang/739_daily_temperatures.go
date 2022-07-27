package main

import "fmt"

// 表示每天的温度，返回一个数组 answer ，
// 其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 方法一：暴力破解，AC。
func dailyTemperatures(temperatures []int) []int {
	var res []int
	for i := 0; i < len(temperatures)-1; i++ {
		flag := false
		for j := i+1; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				res = append(res, j-i)
				flag = true
				break
			}
		}
		if !flag {
			res = append(res, 0)
		}
	}
	res = append(res, 0)
	return res
}

// 方法二：单调栈，AC。
//func dailyTemperatures(temperatures []int) []int {
//	var res []int
//	queue := []int{0}
//	for i := 1; i < len(temperatures); i++ {
//		if temperatures[i] > temperatures[queue[0]] {
//			res = append(res, i - queue[0])
//			if len(queue) > 0 {
//				queue = queue[1:]
//			}
// 		} else {
// 			flag := false
// 			for j := i; j < len(temperatures); j++ {
// 				if temperatures[j] > temperatures[queue[0]] {
//					res = append(res, j - queue[0])
//					flag = true
//					break
//				}
//			}
//			if !flag {
//				res = append(res, 0)
//			}
//			queue = []int{}  // 移除前面所有
//		}
//		queue = append(queue, i)
//	}
//	res = append(res, 0)
//	return res
//}


func main() {
	t := []int{73,74,75,71,69,72,76,73} // [1,1,4,2,1,1,0,0]
	//t := []int{30,40,50,60} // [1,1,1,0]
	//t := []int{30,60,90} // [1,1,0]
	fmt.Println(dailyTemperatures(t))
}