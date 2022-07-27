package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	key int
	value int
}

// 返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案
func topKFrequent(nums []int, k int) []int {
	var res []int
	dict := make(map[int]int)
	for _, v := range nums {
		_, ok := dict[v]
		if ok {
			dict[v] += 1
		} else {
			dict[v] = 1
		}
	}
	var PairList []Pair
	for k, v := range dict {
		PairList = append(PairList, Pair{k, v})
	}
	sort.Slice(PairList, func(i, j int) bool {
		return PairList[i].value > PairList[j].value
	})
	//fmt.Println(PairList)
	for i := 0; i < k; i++ {
		res = append(res, PairList[i].key)
	}
	return res
}

func main() {
	nums, k := []int{1,1,1,2,2,3}, 2
	fmt.Println(topKFrequent(nums, k))
}