package main

// 1. 自己写，未通过
//func hasPathSum(root *TreeNode, targetSum int) bool {
//	if root == nil {
//		return false
//	}
//	if getPathSum(root, 0) == targetSum {
//		return true
//	} else {
//		return false
//	}
//}
//
//func getPathSum(root *TreeNode, i int) int {
//	if root == nil {
//		return 0
//	}
//	v := i + root.Val
//	if root.Left == nil && root.Right == nil {
//		return v
//	}
//	return getPathSum(root.Left, v) + getPathSum(root.Right, v)
//}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}