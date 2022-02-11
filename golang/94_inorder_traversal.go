package main

func inorderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode)
	res := []int{}
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
