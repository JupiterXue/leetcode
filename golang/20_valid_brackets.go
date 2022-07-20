package main

import (
	"fmt"
	_ "strings"
)

func isValid(s string) bool {
	bracketsDict := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	var stack []string
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[i])) // push
			continue
		}
		top := stack[len(stack)-1]
		if bracketsDict[top] == string(s[i]) {
			stack = stack[:len(stack)-1] // pop
		} else {
			stack = append(stack, string(s[i])) // push
		}
		fmt.Println(stack)
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

//func isValid(s string) bool {
//	//for strings.Contains("{}", s) || strings.Contains("[]", s) || strings.Contains("()", s) {
//	for strings.Contains(s, "{}") || strings.Contains(s, "[]") || strings.Contains(s, "()") {
//		s = strings.Replace(s, "{}", "", -1)
//		s = strings.Replace(s, "[]", "", -1)
//		s = strings.Replace(s, "()", "", -1)
//	}
//	return len(s) == 0
//}

//func isValid(s string) bool {
//	bracketsDic := map[string]string{
//		"(" : ")",
//		"[" : "]",
//		"{" : "}",
//	}
//
//	var stack []string
//	stack = append(stack, string(s[0]))
//	for i := 1; i < len(s); i++ {
//		if len(stack) == 0 {
//			stack = append(stack, string(s[i]))
//			continue
//		}
//		if _, ok := bracketsDic[string(s[i])]; ok {
//			stack = append(stack, string(s[i]))
//			continue
//		}
//			preBracket := stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		if bracketsDic[preBracket] == string(s[i]) {
//			continue
//		}
//		return false
//	}
//	if len(stack) != 0 {
//		return false
//	} else {
//		return true
//	}
//}

func main() {
	//s := "{[]}"
	//s := "{}[]"
	//s := "{}["
	s := "(([]){})"
	//s := "()[]{}"
	fmt.Println(isValid(s))
}