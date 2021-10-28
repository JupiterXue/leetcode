package main

import "fmt"

func divide(dividend int, divisor int) int {
	const Maxint32 = int(^uint32((0)) >> 1)
	const Minint32 = ^Maxint32
	//const MaxUint = ^uint(0)
	//const MinUint = 0
	//const MaxInt = int(MaxUint >> 1)
	//const MinInt = -MaxInt - 1
	res := dividend / divisor
	//fmt.Println(Maxint32)
	//fmt.Println(Minint32)
	if res > Maxint32 {
		res--
	} else if res < Minint32 {
		res++
	}
	return res
}

func main() {
	//dividend, divisor := 10, 3
	//dividend, divisor := 7, -3
	dividend, divisor := -2147483648,-1
	fmt.Println(divide(dividend, divisor))
}