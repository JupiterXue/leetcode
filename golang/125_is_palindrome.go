package main

import (
	"fmt"
	"strings"
)

func stringFilter(s string) []string {
	s = strings.ToLower(s)
	res := []string{}
	for _, v := range s {
		if (47 < v && v < 58) || (96 < v && v < 123) {
			res = append(res, string(v))
		}
	}
	return res
}

func isPalindrome2(s string) bool {
	sList := stringFilter(s)
	fmt.Println(sList)
	for i := 0; i < len(sList) / 2; i++ {
		if sList[i] != sList[len(sList)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	//s := "A man, a plan, a canal: Panama"
	//s := "race a car"
	s := "0P"
	fmt.Println(isPalindrome2(s))
}
