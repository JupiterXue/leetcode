package main

import (
	"fmt"
	"math"
)

//func checkPartS(sPart, t string) bool {
//	tmp := []string{}
//	for _, v := range t {
//		tmp = append(tmp, string(v))
//	}
//	fmt.Println(sPart)
//	for _, v := range sPart {
//		for i := 0; i < len(tmp); i++ {
//			if string(v) == tmp[i] {
//				tmp = append(tmp[:i], tmp[i+1:]...)
//				i--
//			}
//		}
//	}
//	fmt.Println(tmp)
//	return len(tmp) == 0
//}
//
//func minWindow(s string, t string) string {
//	winSize := len(t)
//	res := []string{}
//	for i := 0; i <= len(s) - winSize; i++ {
//		for j := i+winSize; j <= len(s); j++ {
//			if checkPartS(s[i:j], t) {
//				res = append(res, s[i:j])
//			}
//		}
//	}
//	//fmt.Println(res)
//	minStrLen := len(s)
//	minIndex := -1
//	for i, v := range res {
//		if minStrLen >= len(v) {
//			minIndex = i
//			minStrLen = len(v)
//		}
//	}
//
//	if minIndex == -1 {
//		return ""
//	} else {
//		return res[minIndex]
//	}
//}

//func minWindow(s string, t string) string {
//	pre, cur := map[byte]int{}, map[byte]int{}
//	for i:= 0; i < len(t); i++ {
//		pre[t[i]]++
//	}
//
//	sLen := len(s)
//	len := math.MaxInt32
//	ansL, ansR := -1, -1
//
//	check := func() bool {
//		for k, v := range pre {
//			if cur[k] < k {
//				return false
//			}
//		}
//		return true
//	}
//
//	for c,r := 0, 0; r < sLen; r++ {
//		if r < sLen && pre[s[r]] > 0 {
//			cur[s[r]]++
//		}
//		for check() && c <= r {
//			if r - c +1 < len {
//				len = r - c + 1
//				ansL, ansR = c, c + len
//			}
//			if _, ok := pre[s[c]]; ok {
//				cur[s[c]]--
//			}
//			c++
//		}
//	}
//	if ansL == -1 {
//		return ""
//	} else {
//		return s[ansL:ansR]
//	}
//}

func minWindow(s string, t string) string {
	var res string
	hasmap := make(map[byte]int)
	left, right := 0, 0
	count := math.MaxInt32

	for i := 0; i < len(t); i++ {
		hasmap[t[i]]++
	}

	for right < len(s) {
		hasmap[s[right]]--
		for checkMap(hasmap) { // 核心部分
			if right - left + 1 < count {
				count = right - left + 1
				res = s[left:right+1]
			}
			hasmap[s[left]]++
			left++
		}
		right++
	}
	return res
}

func checkMap(hasmap map[byte]int) bool {
	for _, v := range hasmap {
		if v > 0 {
			return false
		}
	}
	return true
}

func main() {
	//s := "ADOBECODEBANC"
	//t := "ABC"
	//s := "a"
	//t := "a"
	s := "acbbaca"
	t := "aba"
	fmt.Println(minWindow(s, t))
}
