package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	// 建立两个哈希表，相互进行映射，如果匹配不上就判断不为同构
	sHashTable := map[byte]byte{}
	tHashTable := map[byte]byte{}

	for i := 0; i <len(s); i++ {
		if sHashTable[s[i]] > 0 && sHashTable[s[i]] != t[i] || tHashTable[t[i]] > 0 && tHashTable[t[i]] != s[i] {
			return false
		}
		sHashTable[s[i]] = t[i]
		tHashTable[t[i]] = s[i]
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))
}
