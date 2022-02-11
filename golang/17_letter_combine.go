package main

import (
	"fmt"
	"strings"
)

// 1. 自己写，哈希表，AC
//func letterCombinations(digits string) []string {
//	res := []string{}
//	if len(digits) == 0 {
//		return res
//	}
//	keyboradMap := make(map[string][]string)
//	//keyboradMap["1"] = []string{}
//	keyboradMap["2"] = []string{"a", "b", "c"}
//	keyboradMap["3"] = []string{"d", "e", "f"}
//	keyboradMap["4"] = []string{"g", "h", "i"}
//	keyboradMap["5"] = []string{"j", "k", "l"}
//	keyboradMap["6"] = []string{"m", "n", "o"}
//	keyboradMap["7"] = []string{"p", "q", "r", "s"}
//	keyboradMap["8"] = []string{"t", "u", "v"}
//	keyboradMap["9"] = []string{"w", "x", "y", "z"}
//	for _, v := range digits {
//		num := string(v)
//		for _, letter := range keyboradMap[num] {
//			if len(res) >= len(keyboradMap[num]) {
//				for i := 0; i < len(keyboradMap[num]); i++ {
//					res = append(res, res[i] + letter)
//				}
//			} else {
//				res = append(res, letter)
//			}
//		}
//	}
//	if len(digits) > 1 {
//		return res[len(keyboradMap[string(digits[0])]):]
//	} else {
//		return res
//	}
//}


// 2. 自己写，回溯
func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	letterArray := [10]string{
		"", // 0
		"", // 1
		"abc", // 2
		"def", // 3
		"ghi", // 4
		"jkl", // 5
		"mno", // 6
		"pqrs", // 7
		"tuv", // 8
		"wxyz", // 9
	}
	path := []string{}
	var backTracking func(digits string, index int)
	backTracking = func(digits string, index int, ) {
		if index == len(digits) {
			res = append(res, strings.Join(path, ""))
			return
		}
		//digit, _ := strconv.Atoi(digits[index])
		digit := int(digits[index] - '0')
		letters := letterArray[digit]
		for i := 0; i < len(letters); i++ {
			path = append(path, string(letters[i]))
			backTracking(digits, index+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(digits, 0)
	return res
}

func main() {
	//digits := "23"
	//digits := "2"
	digits := ""
	fmt.Println(letterCombinations(digits))
}
