package main

import "fmt"

// 1. 参考题解，AC！ DFS
//func dfsSymmetric(p, q *TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	}
//
//	if p == nil || q == nil {
//		return false
//	}
//
//	return p.Val == q.Val && dfsSymmetric(p.Left, q.Right) && dfsSymmetric(p.Right, q.Left)
//}

//func isSymmetric(root *TreeNode) bool {
//	return dfsSymmetric(root, root)
//}

// 2. 自己写
func isSymmetric(root *TreeNode) bool {
	return dfsSymmetric(root, root)
}

func dfsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && dfsSymmetric(left.Left, right.Right) && dfsSymmetric(left.Right, right.Left)
}


func main() {
	root := new(TreeNode)
	fmt.Println(isSymmetric(root))
}
