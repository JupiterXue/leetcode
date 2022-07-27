package main

import (
	"fmt"
	"strconv"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}


//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	}
//
//	if p == nil || q == nil {
//		return false
//	}
//
//	if p.Val != q.Val {
//		return false
//	}
//
//	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
//}

func (t *TreeNode) CreateTree(s string)  {
	for i := 0; i < len(s) - 2; i += 3 {
		t.Val, _ = strconv.Atoi(string(s[i]))
		t.Left.Val, _ = strconv.Atoi(string(s[i+1]))
		t.Right.Val, _ = strconv.Atoi(string(s[i+2]))

	}
}

func main() {
	p := new(TreeNode)

	q := new(TreeNode)

	fmt.Println(isSameTree(p, q))
}
