package main

import "fmt"

var depthList []int

func countDepth(root *TreeNode, count int) {
	if root == nil {
		return
	}
	count++
	if root.Left == nil && root.Right == nil {
		depthList = append(depthList, count)
		return
	}
	countDepth(root.Left, count)
	countDepth(root.Right, count)
}

func maxDepth(root *TreeNode) int {
	countDepth(root, 0)
	maxDepthValue := 0
	for _, v := range depthList {
		if maxDepthValue < v {
			maxDepthValue = v
		}
	}
	return maxDepthValue
}

//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	leftValue := maxDepth(root.Left)
//	rightValue := maxDepth(root.Right)
//	if leftValue > rightValue {
//		return leftValue + 1
//	} else {
//		return rightValue + 1
//	}
//}

func main() {
	root := new(TreeNode)
	fmt.Println(maxDepth(root))
}
