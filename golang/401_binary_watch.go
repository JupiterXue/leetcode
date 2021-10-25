package main

import (
	"fmt"
	"math/bits"
)

//func TimeSearch(hours [4]int, mins [6]int, validTime string, res *[]string, shour, sminute, h, m, num int) {
//	if num ==0 {
//		validTime += strconv.Itoa(shour)
//		validTime += ":"
//		if sminute <= 9 {
//			validTime += "0"
//		}
//		validTime += strconv.Itoa(sminute)
//		*res = append(*res, validTime)
//		return
//	}
//
//	// 确认分钟
//	for i := m; i < 6; i++ {
//		if sminute+mins[i] <= 59 {
//			// 回溯回来时，m 之前不再进行选择，因为已经选择并判断过
//			m++
//			TimeSearch(hours, mins, validTime, res, shour, sminute+mins[i], h, m, num-1)
//		}
//	}
//
//	// 确认小时
//	for i := h; i < 4; i++ {
//		if shour + hours[i] <= 11 {
//			h++
//			TimeSearch(hours, mins, validTime, res, shour+hours[i], sminute, h, m, num-1)
//		}
//	}
//}
//
//
//func readBinaryWatch(turnedOn int) []string {
//	hours := [4]int{1,2,4,8}
//	mins := [6]int{1,2,4,8,16,32}
//
//	var result []string
//	validTime := "" // 转换后对应时间
//	shour := 0		// 计算小时
//	smin := 0		// 计算分钟
//
//	if turnedOn == 0 {
//		validTime = "0:00"
//		result = append(result, validTime)
//		return result
//	}
//
//	TimeSearch(hours, mins, validTime, &result, shour, smin, 0, 0, turnedOn)
//
//	return result
//}

// 方法一：枚举时分
/*func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h) + bits.OnesCount8(m) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h ,m))
			}
		}
	}
	return res
}
*/

//方法二：二进制枚举
func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	for i := 0; i < 1024; i++ {
		h, m := i >> 6, i & 63 // 用位运算取出最高4位和最低6位
		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn {
			res = append(res, fmt.Sprintf("%d:%02d", h ,m))
		}
	}
	return res
}


func main() {
	turnedOn := 1
	fmt.Println(readBinaryWatch(turnedOn))
}
