package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		wordsList := strings.Split(sentence, " ")
		fmt.Println(wordsList)
		if len(wordsList) > max {
			max = len(wordsList)
		}
	}
	return max
}

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}
