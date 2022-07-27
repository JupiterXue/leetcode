package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func (this *TreeNode) FrontCreate(nums []int) {
	stack := []*TreeNode{this}
	for _, v := range nums {
		root := stack[len(stack)-1]
		if root == nil {
			continue
		}
		stack = stack[:len(stack)-1]

		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if v == -1 {
			root = nil
		} else {
			root.Val = v
		}
	}
}

func (this *TreeNode) LevelShow() {
	stack := []*TreeNode{this}
	for len(stack) > 0 {
		root := stack[len(stack)-1]
		fmt.Println(root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
}