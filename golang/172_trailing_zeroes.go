package main

import "fmt"

func getTrailNum(n int) int {
	if n <= 1 {
		return n
	}
	return n * getTrailNum(n-1)
}

// 整数溢出
//func trailingZeroes(n int) int {
//	if n == 0 {
//		return 0
//	}
//	trailNum := getTrailNum(n)
//	fmt.Println(trailNum)
//	zero := trailNum % 10
//	count := 0
//	for zero == 0 {
//		count++
//		trailNum /= 10
//		zero = trailNum % 10
//	}
//	return count
//}

func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		count += n/5
		n /= 5
	}
	return count
}


func main() {
	n := 30
	fmt.Println(trailingZeroes(n))
}
