package main

import (
	"fmt"
)

//var depthList []int

//func countDepth(root *TreeNode, count int) {
//	if root == nil {
//		return
//	}
//	count++
//	if root.Left == nil && root.Right == nil {
//		depthList = append(depthList, count)
//		return
//	}
//	countDepth(root.Left, count)
//	countDepth(root.Right, count)
//}
//
//func maxDepth(root *TreeNode) int {
//	countDepth(root, 0)
//	maxDepthValue := 0
//	for _, v := range depthList {
//		if maxDepthValue < v {
//			maxDepthValue = v
//		}
//	}
//	return maxDepthValue
//}

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


// 自己做，没做出来。
//var res []int
//
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	getDepth(root, 0)
//	fmt.Println(res)
//	sort.Ints(res)
//	return res[len(res)-1]
//}
//
//func getDepth(root *TreeNode, depth int){
//	if root == nil {
//		res = append(res, depth)
//		return
//	}
//	if root.Left != nil {
//		getDepth(root.Left, depth+1)
//	}
//	if root.Right != nil {
//		getDepth(root.Right, depth+1)
//	}
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxD(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxD(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func main() {
	//nums := []int{3,9,20,-1,-1,15,7}
	//root := TreeNode{}
	//root.FrontCreate(nums)
	//root.LevelShow()
	root := new(TreeNode)  // new 返回指针
	fmt.Println(maxDepth(root))
}
