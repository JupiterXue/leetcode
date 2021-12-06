package main

import "fmt"

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt2(s string) int {
	// 常规情况可将每个字符当做单独值，累加即可
	// 如果小的数字在大的数字左边，需要减去小的。因此可以取反
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		curInt := romanMap[s[i]]
		nextInt := romanMap[s[i+1]]
		if curInt < nextInt {
			sum -= curInt
		} else {
			sum += curInt
		}
	}
	sum += romanMap[s[len(s)-1]]
	return sum
}

func main() {
	//s := "I"
	s := "LVIII"
	fmt.Println(romanToInt2(s))
}
