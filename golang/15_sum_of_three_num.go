package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 双指针问题
	// 1. 排序
	// 2. 指针1在最前，指针2在最后
	// 3. 类似快排思想
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums) // 学习了
	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for j, k := i+1, len(nums) - 1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				tmp := []int{nums[i],nums[j], nums[k]}
				res = append(res, tmp)
				for j+1 < k && nums[j] == nums[j+1] {
					j++
				}
				for k-1 > j && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum > 0{
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func main() {
	nums := []int{1}
	fmt.Println(threeSum(nums))
}
