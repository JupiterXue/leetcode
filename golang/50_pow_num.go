package main

import (
	"fmt"
)

// 1. 自己写，AC！，用上 math 函数简单粗暴。
//func myPow(x float64, n int) float64 {
//	return math.Pow(x, float64(n))
//}

// 2. 递归
func myPow(x float64, n int) float64 {
	if n < 0 {
		n *= -1
		x = 1/x
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n % 2 > 0 {
		return myPow(x, n-1) * x
	} else {
		return myPow(x*x, n/2)
	}
}


func main() {
	//x := 2.1000
	//n := 3
	//x := 2.000
	//n := 10
	//n := -2
	x := 0.00001
	n := 2147483647
	fmt.Println(myPow(x, n))
}