package main

// 1. 自己写，未通过。
//func isBalanced(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	if checkHigh(root, 1) == -1 {
//		return false
//	} else {
//		return true
//	}
//}
//
//func checkHigh(root *TreeNode, high int) int {
//	if root == nil {
//		return high
//	}
//	left := checkHigh(root.Left, high+1)
//	right := checkHigh(root.Right, high+1)
//	if left - right > 1 || right - left > 1 {
//		return -1
//	}
//	return high
//}

// 2. 官方，自底向上
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return maxHeight(left, right) + 1
}

func maxHeight(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		return -1*x
	}else {
		return x
	}
}