package main

import "fmt"

func lemonadeChange(bills []int) bool {
	// left 表示还剩多少，下标0 是 5元 个数，下标1 是 10元 个数
	left := [2]int{0,0}
	// 第一个元素不为 5元，直接退出
	if bills[0] != 5 {
		return false
	}
	for i := 0; i < len(bills); i++ {
		// 统计 5元 个数
		if bills[i] == 5 {
			left[0]++
		}
		// 统计 10元 个数
		if bills[i] == 10 {
			left[1]++
		}
		// 处理找零情况
		tmp := bills[i] - 5
		if tmp == 5 {
			if left[0] > 0 {
				left[0]--
			} else {
				return false
			}
		}
		if tmp == 15 {
			if left[1] > 0 && left[0] > 0 {
				left[0]--
				left[1]--
			} else if left[1] == 0 && left[0] > 2 {
				left[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	bills := []int{5,5,5,10,20}
	fmt.Println(lemonadeChange(bills))
}
