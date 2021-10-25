package main

import "fmt"

func makeEqual(words []string) bool {
	wordsNum := len(words)
	if wordsNum <= 1 {
		return true
	}
	allWords := ""
	for _, v := range words {
		allWords += v
	}

	charMap := make(map[string]int)
	for _, v := range allWords {
		if _, ok := charMap[string(v)]; ok {
			charMap[string(v)]++
		} else {
			charMap[string(v)] = 1
		}
	}


	fmt.Println(charMap)
	for _, v := range charMap {
		//if v != wordsNum {
		//	return false
		//}
		if v % wordsNum != 0 {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"caaaaa","aaaaaaaaa","a","bbb","bbbbbbbbb","bbb","cc","cccccccccccc","ccccccc","ccccccc","cc","cccc","c","cccccccc","c"}
	fmt.Println(makeEqual(words))
}
