package main

import "fmt"

//func isPowerOfThree(n int) bool {
//	for n > 0 {
//		if n % 3 == 0 {
//			n /= 3
//			continue
//		}
//		if n == 1 {
//			return true
//		} else {
//			return false
//		}
//	}
//	return false
//}

func isPowerOfThree(n int) bool {
	if n == 1{
		return true
	}
	if n == 0 {
		return false
	}
	if n % 3 == 0 {
		return isPowerOfThree(n/3)
	} else {
		return false
	}
}

func main() {
	n := 27
	fmt.Println(isPowerOfThree(n))
}
