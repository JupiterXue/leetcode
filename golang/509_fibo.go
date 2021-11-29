package main

import "fmt"

// 1. 递推，AC！
//func fib(n int) int {
//	f0, f1 := 0, 1
//	for i := 0; i < n; i++ {
//		f0, f1 = f1, f1+f0
//	}
//	return f0
//}

// 2. 递归，AC！
func fib(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	n := 4
	fmt.Println(fib(n))
}
