package main

import "fmt"

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	low := 0
	high := len(s)-1
	for low < high {
		if s[low] != s[high] {
			return 2
		}
		low++
		high--
	}
	return 1
}

func main() {
	s := "ababa"
	fmt.Println(removePalindromeSub(s))
}
