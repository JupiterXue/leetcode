package main

import (
	"fmt"
	"strings"
)

func addSpaces(s string, spaces []int) string {
	lastIdx := 0
	res := []string{}
	for len(spaces) != 0 {
		idx := spaces[0]
		spaces = spaces[1:]
		//tmp := s[lastIdx:idx] + " "
		tmp := s[lastIdx:idx]
		//fmt.Println(tmp)
		res = append(res, tmp)
		lastIdx = idx
	}
	res = append(res, s[lastIdx:])
	return strings.Join(res, " ")
}

func main() {
	//s := "LeetcodeHelpsMeLearn"
	//s := "icodeinpython"
	s := "spacing"
	//spaces := []int{8,13,15}
	//spaces := []int{1,5,7,9}
	spaces := []int{0,1,2,3,4,5,6}
	fmt.Println(addSpaces(s, spaces))
}
