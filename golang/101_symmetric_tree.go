package main

import "fmt"

func dfsSymmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && dfsSymmetric(p.Left, q.Right) && dfsSymmetric(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {
	return dfsSymmetric(root, root)
}

func main() {
	root := new(TreeNode)
	fmt.Println(isSymmetric(root))
}
