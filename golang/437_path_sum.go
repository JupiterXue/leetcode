package main

//var res2 []int
//
//func pathDfs(root *TreeNode, n int) {
//	if root == nil {
//		return
//	}
//	res2 = append(res2, n+root.Val)
//	pathDfs(root.Right, n+root.Val)
//	pathDfs(root.Left, n+root.Val)
//}
//
//func pathSum(root *TreeNode, targetSum int) int {
//	if root == nil {
//		return 0
//	}
//	//res2 = append(res2, root.Val)
//	pathDfs(root.Right, 0)
//	pathDfs(root.Left, 0)
//
//	count := 0
//	fmt.Println(res2)
//	for _, v := range res2 {
//		if v == targetSum {
//			count++
//		}
//	}
//	return count
//}

//func dfsPathSum(root *TreeNode, prefixSumCount map[int][]int, targetsum, currentSum int) int {
//	// map作为函数入参是作为指针进行传递的
//	if root == nil {
//		return 0
//	}
//
//	res := 0
//	currentSum += root.Val
//
//	res += prefixSumCount[currentSum-targetsum][0]
//	prefixSumCount[currentSum] = append(prefixSumCount[currentSum],prefixSumCount[currentSum][0] + 1)
//
//	res += dfsPathSum(root.Left, prefixSumCount, targetsum, currentSum)
//	res += dfsPathSum(root.Right, prefixSumCount, targetsum, currentSum)
//
//	prefixSumCount = append(prefixSumCount[currentSum], prefixSumCount[currentSum][0]-1)
//	return res
//}

//func pathSum(root *TreeNode, targetSum int) int {
//	prefixSumCount := make(map[int][]int)
//	prefixSumCount[0] = append(prefixSumCount[0], 1)
//
//	return dfsPathSum(root, prefixSumCount, targetSum, 0)
//}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return dfsPahtSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}


func dfsPahtSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	tmp := 0
	sum -= root.Val
	if sum == 0 {
		tmp++
	}
	return tmp + dfsPahtSum(root.Left, sum) + dfsPahtSum(root.Right, sum)
}