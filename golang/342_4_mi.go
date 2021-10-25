package main

import "fmt"

func isPowerOfFour(n int) bool {
	//if n == 1 {
	//	return true
	//}

	for n >= 4 {
		if n % 4 == 0 {
			n /= 4
		} else {
			return false
		}
	}

	if n != 1 {
		return false
	} else {
		return true
	}
}

func main() {
	n := 1
	fmt.Println(isPowerOfFour(n))
}
