package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)  // 初始化二维表格
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				// 如果当前字符匹配上了，+1并推送到对角位置
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// 如果当前字符不匹配，选择下方或右侧最大的推送到对角位
				dp[i+1][j+1] = maxLen(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n] // 返回最后的位置，也就是最大的
}

func maxLen(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//t1, t2 := "abcde", "ace"
	//t1, t2 := "abc", "abc"
	//t1, t2 := "bl", "yby"
	t1, t2 := "ezupkr", "ubmrapg"
	fmt.Println(longestCommonSubsequence(t1, t2))
}
