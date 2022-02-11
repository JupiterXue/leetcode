package main

func postorderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode)
	res := []int{}
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}
