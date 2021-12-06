package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	wordMap := map[string]int{}
	wordMap2 := map[string]int{}
	for _, v := range words1 {
		if ok := wordMap[v]; ok > 0 {
			wordMap[v]++
		} else {
			wordMap[v] = 1
		}
	}

	for _, v := range words2 {
		if ok := wordMap2[v]; ok > 0 {
			wordMap2[v]++
		} else {
			wordMap2[v] = 1
		}
	}

	count := 0
	for k := range wordMap {
		if ok := wordMap2[k]; ok > 0 {
			if wordMap[k] == 1 && wordMap2[k] == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	words1 := []string{"leetcode","is","amazing","as","is"}
	words2 := []string{"amazing","leetcode","is"}
	fmt.Println(countWords(words1, words2))
}
