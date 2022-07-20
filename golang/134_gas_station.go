package main

import "fmt"

// 1. 自己写，贪婪，未通过
//func canCompleteCircuit(gas []int, cost []int) int {
//	gasAll := 0
//	num := -1
//	maxGas := 0
//	for i := 0; i < len(gas); i++ {
//		if gas[i] - cost[(i + 1) % len(gas)] > maxGas {
//			num = i
//			maxGas = gas[i] - cost[(i + 1) % len(gas)]
//			gasAll = gas[i]
//		}
//	}
//	//fmt.Println(maxGas)
//	//fmt.Println(num)
//	//fmt.Println()
//	//fmt.Println()
//	//fmt.Println()
//	cur := (num + 1) % len(gas)
//	for cur != num {
//		fmt.Println(gasAll)
//		fmt.Println(cost[cur])
//		fmt.Println()
//		if cost[cur] > gasAll {
//			return -1
//		}
//		gasAll += gas[cur] - cost[cur]
//		//if cur - 1 < 0 {
//		//	gasAll += gas[cur] - cost[len(gas)-1]
//		//} else {
//		//	gasAll += gas[cur] - cost[(cur - 1) % len(gas)]
//		//}
//		cur = (cur + 1) % len(gas)
//	}
//	return num
//}


// 2. 随想录，贪婪，AC！
func canCompleteCircuit(gas []int, cost []int) int {
	curSum, totalSum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}

func main() {
	//gas, cost := []int{1,2,3,4,5}, []int{3,4,5,1,2}
	//gas, cost := []int{2, 3, 4}, []int{3, 4, 3}
	//gas, cost := []int{5,1,2,3,4}, []int{4,4,1,5,1}
	gas, cost := []int{1,2,3,4,5}, []int{3,4,5,1,2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
