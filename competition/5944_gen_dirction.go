package main

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func searchVal(root *TreeNode, startValue, destValue int, res string) string {
	if root.Val == startValue {
		return "S"
	}
	if root.Val == destValue {
		return "E"
	}
	r1 := ""
	if root.Left == nil {
		res += "L"
		r1 = searchVal(root.Left, startValue, destValue, res)
	}
	r2 := ""
	if root.Right == nil {
		res += "R"
		r2 = searchVal(root.Right, startValue, destValue, res)
	}
	return r1 + r2
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	res := "U"
	res = searchVal(root, startValue, destValue, res)
	resList := []string{}
	isRoot := false
	for _, v := range res {
		char := string(v)
		if char == "U" {
			if isRoot {
				continue
			}
			isRoot = true
		}
		if isRoot {
			resList = append(resList, "U")
		} else {
			resList = append(resList, char)
		}
	}
	return res
}