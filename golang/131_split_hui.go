package main

import (
	"fmt"
	"strings"
)

func partition3(s string) [][]string {
	var res [][]string
	var pathList []string
	var isPalindrome func(s string) bool
	var backTracking func(k string, startIndex int, pathList []string)
	isPalindrome = func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-i-1] {
				return false
			}
		}
		return true
	}

	backTracking = func(k string, startIndex int, pathList []string) {
		if len(strings.Join(pathList, "")) > len(k) {
			return
		} else if strings.Join(pathList, "") == k {
			res = append(res, pathList)
		}
		for i := startIndex; i < len(k); i++ {
			pathList = append(pathList, k[startIndex:i+1])
			backTracking(k, i+1, pathList)
			pathList = pathList[:len(pathList)-1]
		}
	}
	backTracking(s, 0, pathList)
	//fmt.Println(res)
	for i := 0; i < len(res); i++ {
		for _, path := range res[i] {
			if !isPalindrome(path) {
				res = append(res[:i], res[i+1:]...)
				i--
			}
		}
	}
	return res
}

func main() {
	//s := "aab"
	//s := "a"
	s := "abbab"
	fmt.Println(partition3(s))
}
