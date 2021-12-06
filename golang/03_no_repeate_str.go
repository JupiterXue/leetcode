package main

import "fmt"

// 1. 自己写，AC！
//func check(partS string) bool {
//	for i := 0; i < len(partS) - 1; i++ {
//		fmt.Println(string(partS[i]))
//		fmt.Println(string(partS[i+1]))
//		fmt.Println()
//		if string(partS[i]) == string(partS[i+1]) {
//			return false
//		}
//	}
//	return true
//}
//
//func lengthOfLongestSubstring(s string) int {
//	OrderLen := 0
//	for i := 0; i < len(s); i++ {
//		if check(s[i:]) {
//			OrderLen = len(s[i:])
//			break
//		}
//	}
//
//	ROrderLen := 0
//	for i := 0; i < len(s); i++ {
//		if check(s[:len(s)-i]) {
//			ROrderLen = len(s[:len(s)-i])
//			break
//		}
//	}
//
//	fmt.Println(OrderLen)
//	fmt.Println(ROrderLen)
//
//	if OrderLen >= ROrderLen {
//		return OrderLen
//	} else {
//		return ROrderLen
//	}
//}


// 2. 自己写，哈希表，超时。。。
//func lengthOfLongestSubstring(s string) int {
//	windowsSize := len(s)
//	for windowsSize > 0 {
//		start := 0
//		for i := windowsSize; i <= len(s); i++ {
//			fmt.Println(s[start:i])
//			if hashCheck(s[start:i]) {
//				return len(s[start:i])
//			}
//			start++
//		}
//		windowsSize--
//	}
//	return 0
//}
//
//func hashCheck(s string) bool {
//	tmpMap := map[string]int{}
//	for _, v := range s {
//		if ok := tmpMap[string(v)]; ok > 0 {
//			return false
//		} else {
//			tmpMap[string(v)] = 1
//		}
//	}
//	return true
//}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)

	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk + 1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}


func main() {
	//s := "abcabcbb"
	//s :=  "bbbbb"
	//s :=  "pwwkew"
	//s :=  ""
	//s :=  " "
	//s :=  "au"
	//s :=  "aa"
	s := "mrjkdfwfsfjoblbhtjcpdbjdqkvevshhjssnzosstdgwqhelqibumkzcwujsnsbyktlkkgeflkectkpjuqfgdgjbduvqm"
	fmt.Println(lengthOfLongestSubstring(s))
}