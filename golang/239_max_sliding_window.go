package main

import (
	"fmt"
)

// 有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回滑动窗口中的最大值 。

// 思路1：两重循环迭代窗口找最大值，超时。
//func maxSlidingWindow(nums []int, k int) []int {
//	var res []int
//	winIdx := k - 1
//	for winIdx < len(nums) {
//		tmp := -10000
//		for i := winIdx - k + 1; i <= winIdx; i++ {
//			if nums[i] > tmp {
//				tmp = nums[i]
//			}
//		}
//		res = append(res, tmp)
//		winIdx += 1
//	}
//	return res
//}


// 思路2：没做出来，参考题解，单调队列+双指针
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var queue []int
	optPush := func(i int) {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	for i := 0; i < k; i++ {
		optPush(i)
	}
	res = append(res, nums[queue[0]])
	for i := k; i < len(nums); i++ {
		optPush(i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		res = append(res, nums[queue[0]])
	}
	return res
}

func main() {
	//nums, k := []int{1,3,-1,-3,5,3,6,7}, 3  // [3,3,5,5,6,7]
	nums, k := []int{1}, 1  // [1]
	fmt.Println(maxSlidingWindow(nums, k))
}