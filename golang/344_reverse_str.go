package main

import "fmt"

// 1. 递推，AC！
//func reverseString(s []byte) {
//	for i := 0; i < len(s) / 2; i++ {
//		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
//	}
//}

// 2. 递归，AC！
func reverseString(s []byte) {
	i := 0
	n := len(s) / 2
	recursiveStr(i, n, s)
}

func recursiveStr(i, n int, s []byte) {
	if i >= n {
		return
	} else {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
		i++
		recursiveStr(i, n, s)
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	//s := []byte{'h', 'e', 'l', 'f', 'o', 'e'}
	fmt.Println(s)
	reverseString(s)
	fmt.Println(s)
}
