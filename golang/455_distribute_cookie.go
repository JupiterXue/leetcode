package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	count := 0
	sort.Ints(g)
	sort.Ints(s)
	//for i := len(g) - 1; i >= 0; i-- {
	//	for j := len(s) - 1; j >= 0; j-- {
	//		if g[i] <= s[j] {
	//			g = append(g[:i], g[i+1:]...)
	//			s = append(g[:j], g[j+1:]...)
	//		}
	//	}
	//}
	n, m := len(g), len(s)
	for i, j := 0, 0; i < n && j < m; i++ {
		for j < m && g[i] > s[j] {
			j++
		}
		if j < m {
			count++
			j++
		}
	}
	return count
}

func main() {
	g := []int{1,2,3}
	s := []int{1,1}
	fmt.Println(findContentChildren(g, s))
}
