package main

import "fmt"

// 1. 递推, AC
//func isPowerOfTwo(n int) bool {
//	if n <= 0 {
//		return false
//	}
//	for n % 2 == 0 {
//		n /= 2
//	}
//	if n <= 1 {
//		return true
//	} else {
//		return false
//	}
//}

// 2. 递归, AC
//func isPowerOfTwo(n int) bool {
//	if n <= 0 {
//		return false
//	} else if n <= 1 {
//		return true
//	} else {
//		if n % 2 == 0 {
//			return isPowerOfTwo(n/2)
//		} else {
//			return false
//		}
//	}
//}

// 3. 位运算
func isPowerOfTwo(n int) bool {
	return n > 0 && n & (n-1) == 0
}


func main() {
	n := 16
	fmt.Println(isPowerOfTwo(n))
}
