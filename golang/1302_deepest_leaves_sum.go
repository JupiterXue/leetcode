package main

import "fmt"

func deepestLeavesSum(root *TreeNode) int {
	nodeList := []*TreeNode{root}
	res := 0
	for len(nodeList) > 0 {
		n := len(nodeList)
		res = 0
		for i := 0; i < n; i++ {
			cur := nodeList[0]
			nodeList = nodeList[1:]
			res += cur.Val
			if cur.Left != nil {
				nodeList = append(nodeList, cur.Left)
			}
			if cur.Right != nil {
				nodeList = append(nodeList, cur.Right)
			}
		}
	}
	return res
}


//var res = 0
//
//func deepestLeavesSum(root *TreeNode) int {
//	if root.Left == nil && root.Right == nil {
//		res += root.Val
//	}
//
//	if root.Left != nil {
//		deepestLeavesSum(root.Left)
//	}
//
//	if root.Right != nil {
//		deepestLeavesSum(root.Right)
//	}
//
//	return res
//}

func main() {
	//root := []int{1,2,3,4,5,nil,6,7,nil,nil,nil,nil,8}
	root := new(TreeNode)
	fmt.Println(deepestLeavesSum(root))
}
