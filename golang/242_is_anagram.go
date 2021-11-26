package main

import "fmt"

func genDict(s string) map[string]int {
	tmpDict := make(map[string]int)
	for _, v := range s {
		if ok := tmpDict[string(v)]; ok > 0 {
			tmpDict[string(v)]++
		} else {
			tmpDict[string(v)] = 1
		}
	}
	return tmpDict
}

func isAnagram(s string, t string) bool {
	sDict := genDict(s)
	tDict := genDict(t)
	//fmt.Println(sDict)
	//fmt.Println(tDict)
	if len(sDict) != len(tDict) {
		return false
	}
	for k := range sDict {
		fmt.Println(k)
		if sDict[k] != tDict[k] {
			return false
		}
	}
	return true
}

func main() {
	//s := "anagram"
	//t := "nagaram"
	s := "a"
	t := "ab"
	fmt.Println(isAnagram(s, t))
}
