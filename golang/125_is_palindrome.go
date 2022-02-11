package main

import (
	"fmt"
	"strings"
)

//func stringFilter(s string) []string {
//	s = strings.ToLower(s)
//	res := []string{}
//	for _, v := range s {
//		if (47 < v && v < 58) || (96 < v && v < 123) {
//			res = append(res, string(v))
//		}
//	}
//	return res
//}
//
//func isPalindrome2(s string) bool {
//	sList := stringFilter(s)
//	fmt.Println(sList)
//	for i := 0; i < len(sList) / 2; i++ {
//		if sList[i] != sList[len(sList)-1-i] {
//			return false
//		}
//	}
//	return true
//}

// 2. 自己写
func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	var newS string
	for _, v := range s {
		// 判断是否为数字，		 判断是否为小写字母
		if (47 < v && v < 58) || (96 < v && v < 123) {
			newS += string(v)
		}
	}
	s = newS
	//fmt.Println(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	//s := "race a car"
	//s := "0P"
	fmt.Println(isPalindrome2(s))
}
