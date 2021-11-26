package main

import "fmt"

// 1. 自己写，Fail
//func isMatch(s string, p string) bool {
//	// 实现一个支持 '.' 和 '*' 的正则表达式
//	// 字符串 s 和 字符规律 p
//	pList := []string{}
//	start := false
//	for _, v := range p {
//		pChar := string(v)
//		if start {
//			pList = append(pList, pChar)
//		} else {
//			if  pChar == "." || pChar == "*" || pChar == string(s[0]) {
//				start = true
//				pList = append(pList, pChar)
//			}
//		}
//	}
//	isAsterisk := false
//	for _, v := range s {
//		if len(pList) == 0 {
//			return false
//		}
//		sChar := string(v)
//		pChar := pList[0]
//		//fmt.Println(sChar)
//		//fmt.Println(pChar)
//
//		if sChar != pChar {
//			if pChar == "." || isAsterisk {
//				continue
//			} else if pChar == "*" {  // 不用出栈
//				isAsterisk = true
//			} else {
//				return false
//			}
//		} else {
//			isAsterisk = false
//			pList = pList[1:]
//		}
//	}
//	return true
//}

// 2. 动态规划，题解，AC！
//func isMatch(s string, p string) bool {
//	m, n := len(s), len(p)
//	matches := func(i, j int) bool {
//		if i == 0 {
//			return false
//		}
//		if p[j-1] == '.' {
//			return true
//		}
//		return s[i-1] == p[j-1]
//	}
//
//	f := make([][]bool, m + 1)
//	for i := 0; i < len(f); i++ {
//		f[i] = make([]bool, n + 1)
//	}
//	f[0][0] = true
//	for i := 0; i <= m; i++ {
//		for j := 1; j <= n; j++ {
//			if p[j-1] == '*' {
//				f[i][j] = f[i][j] || f[i][j-2]
//				if matches(i, j-1) {
//					f[i][j] = f[i][j] || f[i-1][j]
//				}
//			} else if matches(i, j) {
//				f[i][j] = f[i][j] || f[i-1][j-1]
//			}
//		}
//	}
//	return f[m][n]
//}

// 3. 递归，题解，AC！
func isMatch(s string, p string) bool {
	pLen := len(p)
	sLen := len(s)
	if pLen == 0 && sLen == 0 {
		return true
	}

	if pLen == 0 {
		return false
	}

	if p == ".*" {
		return true
	}
	return matchRecursive(0, 0, s, p)
}

func matchRecursive(si, pi int, s, p string) bool {
	pLen := len(p)
	sLen := len(s)
	// 同时到达尾端，唯一可能返回 true 的情况
	if pi >= pLen && si >= sLen {
		return true
	}
	// 匹配器没有了，但字符还在
	if pi >= pLen && si < sLen {
		return false
	}
	// 分情况向下递归。星号匹配，点号匹配，精准匹配
	if pi + 1 < len(p) && p[pi+1] == byte('*') {
		mateByte := p[pi]
		// 匹配超出，但星号匹配可以忽略并继续进入下个匹配
		if si >= sLen {
			return matchRecursive(si, pi+2, s, p)
		}
		// 匹配到了三种选择。* 特殊，可以匹配任意字符
		// 三种步进方式都有可能
		if s[si] == mateByte || mateByte == byte('.') {
			return matchRecursive(si, pi+2, s, p) || matchRecursive(si+1, pi, s, p) || matchRecursive(si+1, pi+2, s, p)
		}
		return matchRecursive(si, pi+2, s, p)
	} else if p[pi] == byte('.') {
		// 匹配超出
		if si >= sLen {
			return false
		}
		// . 可匹配任何字符，直接双步进
		return matchRecursive(si+1, pi+1, s, p)
	} else {
		// 匹配超出
		if si >= sLen {
			return false
		}
		// 具体字符，严格匹配
		if p[pi] != s[si]{
			return false
		}
		return matchRecursive(si+1, pi+1, s, p)
	}
}

func main() {
	//s := "aab"  // true
	//p := ".*"
	//s := "mississippi"  // true
	//p := "mis*is*ip*."
	s := "mississippi"
	p := "mis*is*p*."
	fmt.Println(isMatch(s, p))
}
