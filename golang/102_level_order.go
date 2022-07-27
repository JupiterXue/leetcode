package main

// 层序遍历，代码随想录，AC
//func levelOrder(root *TreeNode) [][]int {
//	var res [][]int
//	if root == nil {
//		return res
//	}
//	queue := list.New()
//	queue.PushBack(root)
//	var tmp []int
//	for queue.Len() > 0 {
//		n := queue.Len()
//		for i := 0; i < n; i++ {
//			node := queue.Remove(queue.Front()).(* TreeNode)
//			if node.Left != nil {
//				queue.PushBack(node.Left)
//			}
//			if node.Right != nil {
//				queue.PushBack(node.Right)
//			}
//			tmp = append(tmp, node.Val)
//		}
//		res = append(res, tmp)
//		tmp = []int{}
//	}
//	return res
//}

// 题解，栈
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		n := len(stack)
		var v []int
		for i := 0; i < n; i++ {
			node := stack[0]
			stack = stack[1:]
			if node != nil {
				v = append(v, node.Val)
				if node.Left != nil {
					stack = append(stack, node.Left)
				}
				if node.Right != nil {
					stack = append(stack, node.Right)
				}
			}
		}
		if len(v) > 0 {
			res = append(res, v)
		}
	}
	return res
}


func main() {
	// 输入：root = [3,9,20,null,null,15,7]
	// 输出：[[3],[9,20],[15,7]]
	nums := []int{}
	tree := TreeNode{}
	tree.FrontCreate(nums)
	tree.LevelShow()
}