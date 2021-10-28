package main

import "fmt"

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	keyboradMap := make(map[string][]string)
	//keyboradMap["1"] = []string{}
	keyboradMap["2"] = []string{"a", "b", "c"}
	keyboradMap["3"] = []string{"d", "e", "f"}
	keyboradMap["4"] = []string{"g", "h", "i"}
	keyboradMap["5"] = []string{"j", "k", "l"}
	keyboradMap["6"] = []string{"m", "n", "o"}
	keyboradMap["7"] = []string{"p", "q", "r", "s"}
	keyboradMap["8"] = []string{"t", "u", "v"}
	keyboradMap["9"] = []string{"w", "x", "y", "z"}
	for _, v := range digits {
		num := string(v)
		for _, letter := range keyboradMap[num] {
			if len(res) >= len(keyboradMap[num]) {
				for i := 0; i < len(keyboradMap[num]); i++ {
					res = append(res, res[i] + letter)
				}
			} else {
				res = append(res, letter)
			}
		}
	}
	if len(digits) > 1 {
		return res[len(keyboradMap[string(digits[0])]):]
	} else {
		return res
	}
}

func main() {
	//digits := "23"
	digits := "2"
	fmt.Println(letterCombinations(digits))
}
