package main

import "fmt"

func check(partS string) bool {
	for i := 0; i < len(partS) - 1; i++ {
		fmt.Println(string(partS[i]))
		fmt.Println(string(partS[i+1]))
		fmt.Println()
		if string(partS[i]) == string(partS[i+1]) {
			return false
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	OrderLen := 0
	for i := 0; i < len(s); i++ {
		if check(s[i:]) {
			OrderLen = len(s[i:])
			break
		}
	}

	ROrderLen := 0
	for i := 0; i < len(s); i++ {
		if check(s[:len(s)-i]) {
			ROrderLen = len(s[:len(s)-i])
			break
		}
	}

	fmt.Println(OrderLen)
	fmt.Println(ROrderLen)

	if OrderLen >= ROrderLen {
		return OrderLen
	} else {
		return ROrderLen
	}
}


func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}