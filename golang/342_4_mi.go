package main

import "fmt"

// 1. 推导法，AC
//func isPowerOfFour(n int) bool {
//	//if n == 1 {
//	//	return true
//	//}
//
//	for n >= 4 {
//		if n % 4 == 0 {
//			n /= 4
//		} else {
//			return false
//		}
//	}
//
//	if n != 1 {
//		return false
//	} else {
//		return true
//	}
//}

// 2. 递归法，AC
func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	if n == 4 {
		return true
	} else if n < 4 {
		return false
	} else {
		if n % 4 == 0 {
			return isPowerOfFour(n/4)
		} else {
			return false
		}
	}
}

func main() {
	n := 5
	fmt.Println(isPowerOfFour(n))
}
