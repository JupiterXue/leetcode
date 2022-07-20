package main

import (
	"fmt"
	"strings"
)

func removeDuplicates2(s string) string {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[i]))
			continue
		}
		top := stack[len(stack)-1]
		if top == string(s[i]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(s[i]))
		}
		//fmt.Println(stack)
	}
	return strings.Join(stack, "")
}

func main() {
	s := "abbaca"
	fmt.Println(removeDuplicates2(s))	//ca
}
