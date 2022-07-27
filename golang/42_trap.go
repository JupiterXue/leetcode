package main

import "fmt"


// 双指针，随想录，AC
//func trap(height []int) int {
//	var left, right, leftMax, rightMax, res int
//	right = len(height) - 1
//	for left < right {
//		if height[left] < height[right] {
//			if height[left] >= leftMax {
//				leftMax = height[left] // 设置左边最高柱子
//			} else {
//				res += leftMax - height[left] // 右边有柱子挡水，如遇≤ leftMax 的，全部加入水池
//			}
//			left++
//		} else {
//			if height[right] > rightMax {
//				rightMax = height[right] // 设置右边最高柱子
//			} else {
//				res += rightMax - height[right] // 左边有柱子，如遇 ≤ rightMax 的，全部加入水池
//			}
//			right--
//		}
//	}
//	return res
//}


func trap(height []int) int {
	var res int
	var stack []int
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := minHeight(height[left], h) - height[top]
			res += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return res
}

func minHeight(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//h := []int{0,1,0,2,1,0,1,3,2,1,2,1} // 6
	h := []int{4,2,0,3,2,5} // 9
	fmt.Println(trap(h))
}