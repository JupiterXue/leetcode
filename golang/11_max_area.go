package main

import (
	"fmt"
)

// 1. 题解，贪婪，AC！
//func max_num(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
//
//func maxArea(height []int) int {
//	i, j, res := 0, len(height)-1, 0
//	for i < j {
//		if height[i] < height[j] {
//			res = max_num(res, height[i] * (j-i))
//			if res == 17 {
//				fmt.Println(i)
//				fmt.Println(j)
//			}
//			i++
//		} else {
//			res = max_num(res, height[j] * (j-i))
//			if res == 17 {
//				fmt.Println(i)
//				fmt.Println(j)
//			}
//			j--
//		}
//	}
//	return res
//}

// 2. 自己写，贪婪，failed...
//func maxArea(height []int) int {
//	// 1. 距离最大
//	// 2. 高度最大
//	// 3. 高度大，距离大
//	i, j := 0, len(height)-1
//	width := getMin(height[i], height[j])
//	max := width*(j-i)
//	max = getMax(0, max)
//	if len(height) == 2 {
//		return max
//	}
//	for i <= j {
//		i++
//		width2 := getMin(height[i], height[j])
//		max2 := width2*(j-i)
//		max = getMax(max, max2)
//		j--
//		width3 := getMin(height[i], height[j])
//		max3 := width3*(j-i)
//		max = getMax(max, max3)
//	}
//	//mid := len(height)/2
//	for low := 0; low <= len(height); low++ {
//		for high := len(height)-1; high >= low; high--{
//			width := getMin(height[low], height[high])
//			max2 := width*(high-low)
//			max = getMax(max, max2)
//		}
//	}
//	return max
//}

//func getMin(i int, i2 int) int {
//	if i > i2 {
//		return i2
//	} else {
//		return i
//	}
//}
//
//func getMax(i, i2 int) int {
//	if i > i2 {
//		return i
//	} else {
//		return i2
//	}
//}


// 3. 题解，贪心，AC！
func maxArea(height []int) int {
	i, j, res := 0, len(height)-1, 0
	for i < j {
		// 计算当前最大面积
		cur := (j-i)*min2(height[i], height[j])
		if cur > res {
			res = cur
		}
		// 移动较小侧指针，原来是这样
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func min2(i int, i2 int) int {
	if i > i2 {
		return i2
	} else {
		return i
	}
}





func main() {
	//height := []int{1,8,6,2,5,4,8,3,7}
	//height := []int{1,1}
	//height := []int{4,3,2,1,4}
	//height := []int{1,2,4,3}
	//height := []int{2,3,10,5,7,8,9} //36
	height := []int{2,3,4,5,18,17,6} // 17
	fmt.Println(maxArea(height))
}
