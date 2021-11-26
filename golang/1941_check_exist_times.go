package main

import (
	"fmt"
)

func areOccurrencesEqual(s string) bool {
	dict := make(map[string]int)
	for _, c := range s {
		if _, ok := dict[string(c)]; ok  {
			dict[string(c)]++
		} else {
			dict[string(c)] = 1
		}
	}
	n := dict[string(s[0])]
	for _, v := range dict {
		if n != v {
			return false
		}
	}
	return true
}

func main() {
	s := "abacbc"
	fmt.Println(areOccurrencesEqual(s))
}
