package main

import (
	"fmt"
)

func checkPartS(sPart, t string) bool {
	tmp := []string{}
	for _, v := range t {
		tmp = append(tmp, string(v))
	}
	fmt.Println(sPart)
	for _, v := range sPart {
		for i := 0; i < len(tmp); i++ {
			if string(v) == tmp[i] {
				tmp = append(tmp[:i], tmp[i+1:]...)
				i--
			}
		}
	}
	fmt.Println(tmp)
	return len(tmp) == 0
}

func minWindow(s string, t string) string {
	winSize := len(t)
	res := []string{}
	for i := 0; i <= len(s) - winSize; i++ {
		for j := i+winSize; j <= len(s); j++ {
			if checkPartS(s[i:j], t) {
				res = append(res, s[i:j])
			}
		}
	}
	//fmt.Println(res)
	minStrLen := len(s)
	minIndex := -1
	for i, v := range res {
		if minStrLen >= len(v) {
			minIndex = i
			minStrLen = len(v)
		}
	}

	if minIndex == -1 {
		return ""
	} else {
		return res[minIndex]
	}
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
