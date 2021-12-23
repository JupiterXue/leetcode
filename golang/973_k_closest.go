package main

import "fmt"

func kClosest(points [][]int, k int) [][]int {
	nums := []int{}
	for _, line := range points {
		distance := getDistance(line)
		nums = append(nums, distance)
	}
	res := [][]int{}
	for i := 0; i < k; i++ {
		idx := findMinIdx(nums)
		res = append(res, points[idx])
		points = append(points[:idx], points[idx+1:]...)
		nums = append(nums[:idx], nums[idx+1:]...)
	}
	return res
}

func getDistance(nums []int) int {
	return nums[0]*nums[0] + nums[1]*nums[1]
}

func findMinIdx(nums []int) int {
	min := nums[0]
	idx := 0
	for i, v := range nums {
		if min > v {
			min = v
			idx = i
		}
	}
	return idx
}

func main() {
	points := [][]int{{1,3},{-2,2}}
	k := 1
	fmt.Println(kClosest(points, k))
}
