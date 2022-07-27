package main

import "math"

// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minNum := math.MaxInt32
	if root.Left != nil {
		minNum =  minD(minDepth(root.Left), minNum)
	}
	if root.Right != nil {
		minNum =  minD(minDepth(root.Right), minNum)
	}
	return minNum+1
}

func minD(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 3,9,20,null,null,15,7
// res:2

// 2,null,3,null,4,null,5,null,6
// res:5