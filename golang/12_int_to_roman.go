package main

import (
	"fmt"
)



// 1. 自己做，没做出来。
//var intMap = map[int]byte{
//	1: 'I',
//	5: 'V',
//	10: 'X',
//	50: 'L',
//	100: 'C',
//	500: 'D',
//	1000: 'M',
//}
//
//func intToRoman(num int) string {
//	// 注意题目提醒的六种情况，单独做处理
//	size := 1000
//	sizeList := []int{1,5,10,100,500,1000}
//	sizeIndex := 5
//	res := []string{}
//	for i := num; i > 0; i -= size {
//		cur := i / size
//		if cur*size == size - sizeList[sizeIndex-1] ||  cur*size == size - sizeList[sizeIndex-2]{
//			res = append(res, string(intMap[size]))
//			res = append(res, string(intMap[size+1]))
//			i /= 10
//		} else {
//			for j := 0; j < cur; j++ {
//				res = append(res, string(intMap[size]))
//				i -= size
//			}
//		}
//		if sizeIndex > 0 {
//			if i - size < 0 {
//				sizeIndex--
//				size = sizeList[sizeIndex]
//			}
//		}
//	}
//	fmt.Println(res)
//	return strings.Join(res, "")
//}

var valueSymbols = []struct{
	value 	int
	symbol 	string
} {
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}


func main() {
	//num := 3
	//num := 58
	num := 1994
	fmt.Println(intToRoman(num))
}
