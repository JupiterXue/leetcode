package main

//func invertTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//
//	left := invertTree(root.Left)
//	right := invertTree(root.Right)
//	root.Left = right
//	root.Right = left
//	return root
//}

//func invertTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	left := invertTree(root.Left)
//	right := invertTree(root.Right)
//
//	root.Left = right
//	root.Right = left
//	return root
//}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}