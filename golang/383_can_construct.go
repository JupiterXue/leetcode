package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	noteMap := make(map[string]int)
	magMap := make(map[string]int)
	for _, v := range ransomNote {
		if _, ok := noteMap[string(v)]; ok {
			noteMap[string(v)]++
		} else {
			noteMap[string(v)] = 1
		}
	}

	for _, v := range magazine {
		if _, ok := magMap[string(v)]; ok {
			magMap[string(v)]++
		} else {
			magMap[string(v)] = 1
		}
	}

	for k := range noteMap {
		if noteMap[k] > magMap[k] {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "a"
	magazine := "b"
	fmt.Println(canConstruct(ransomNote, magazine))
}
