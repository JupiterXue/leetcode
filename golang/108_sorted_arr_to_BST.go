package main

import "fmt"

//func sortedArrayToBST(nums []int) *TreeNode {
//	// 数组转平衡二叉树。需要复习：二叉搜索树、平衡二叉树、递归
//	// 中点左右相差不大于 1，因此首先要找中点
//	if len(nums) < 1 {
//		return nil
//	}
//	mid := (len(nums) - 1) / 2
//	root := new(TreeNode)
//	root.Val = nums[mid]
//	root.Left = sortedArrayToBST(nums[:mid])
//	root.Right = sortedArrayToBST(nums[mid+1:])
//	return root
//}

//func sortedArrayToBST(nums []int) *TreeNode {
//	if len(nums) < 1 {
//		return nil
//	}
//
//	mid := (len(nums) - 1) / 2
//	node := new(TreeNode)
//	node.Val = nums[mid]
//	node.Left = sortedArrayToBST(nums[:mid])
//	node.Right = sortedArrayToBST(nums[mid+1:])
//	return node
//}



func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	mid := (len(nums)-1)/2
	root := new(TreeNode)
	root.Val = nums[mid]
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func main() {
	nums := []int{-10,-3,0,5,9}
	fmt.Println(sortedArrayToBST(nums))
}