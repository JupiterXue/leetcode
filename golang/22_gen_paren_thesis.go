package main

import (
	"fmt"
)

func gen(left, right int, vret *[]string, str string) {
	if left > right {
		return
	}
	if right == 0 && left == 0 {
		*vret = append(*vret, str)
	}
	if left > 0 {
		gen(left-1, right, vret, str+"(")
	}
	if right > 0 {
		gen(left, right-1, vret, str+")")
	}
}

func generateParenthesis(n int) []string {
	vret := make([]string, 0)
	gen(n ,n, &vret, "")
	return vret
}

//func valid(tmpS []string) bool {
//	count := 0
//	for _, s := range tmpS {
//		if string(s) == "(" {
//			count++
//		} else {
//			count--
//		}
//		if count < 0 {
//			return false
//		}
//	}
//	return count == 0
//}
//
//func gen(res, tmpS []string, n int) {
//	if len(tmpS) == 2*n {
//		if valid(tmpS) {
//			res = append(res, strings.Join(tmpS, ""))
//		} else {
//			tmpS = append(tmpS, "(")
//			gen(res, tmpS, n)
//			tmpS = tmpS[:len(tmpS)-1]
//			tmpS = append(tmpS, ")")
//			tmpS = tmpS[:len(tmpS)-1]
//		}
//	}
//}
//
//func generateParenthesis(n int) []string {
//	res := []string{}
//	gen(res, []string{}, n)
//
//	return res
//}



func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}