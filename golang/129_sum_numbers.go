package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//func sumNumbers(root *TreeNode) int {
//	if root.Left == nil && root.Right == nil {
//		return root.Val
//	}
//
//	return 2 * 10 * root.Val + sumNumbers(root.Left) + sumNumbers(root.Right)
//}

var result = 0

func dfs(node *TreeNode, num int) {
	num = num * 10 + node.Val

	if node.Left == nil && node.Right == nil {
		result += num
	}
	if node.Left != nil {
		dfs(node.Left, num)
	}
	if node.Right != nil {
		dfs(node.Right, num)
	}
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return result
	}
	dfs(root, root.Val)
	return result
}

func main() {
	fmt.Println(sum)
}
