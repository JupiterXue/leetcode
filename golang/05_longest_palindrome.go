package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	// dp[i][j] 表示 s[i.j] 是否是回文串
	dp := make([][]bool, len(s), )

	res := ""


	return res
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
