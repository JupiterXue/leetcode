package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}
	left := searchBST(root.Left, val)
	if left != nil {
		return left
	}
	right := searchBST(root.Right, val)
	if right != nil {
		return right
	}
	return nil
}