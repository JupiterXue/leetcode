package main

// 1. 自己写，未通过！
var resSlice []int

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	resSlice = append(resSlice, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return resSlice
}

//func preorderTraversal(root *TreeNode) []int {
//	var traversal func(node *TreeNode)
//	res := []int{}
//	traversal = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		res = append(res, node.Val)
//		traversal(node.Left)
//		traversal(node.Right)
//	}
//	traversal(root)
//	return res
//}
