package main

import (
	"fmt"
)

//func numEquivDominoPairs(dominoes [][]int) int {
//	count := 0
//	for i := 0; i < len(dominoes)-1; i++ {
//		for j := i+1; j < len(dominoes); j++ {
//			sort.Ints(dominoes[i])
//			sort.Ints(dominoes[j])
//			if reflect.DeepEqual(dominoes[i], dominoes[j]) {
//				count++
//			}
//		}
//	}
//	return count
//}

func numEquivDominoPairs(dominoes [][]int) int {
	cnt := [100]int{}
	count := 0
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1]
		count += cnt[v]
		cnt[v]++
	}
	return count
}


func main() {
	dominoes := [][]int{{1,2},{2,1},{3,4},{5,6}}
	fmt.Println(numEquivDominoPairs(dominoes))
}
