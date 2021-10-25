package main

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) []int {
	res := []int{}
	zeroList := []int{}
	for i, v := range s {
		if byte(v) == c {
			zeroList = append(zeroList, i)
			res = append(res, 0)
		} else {
			res = append(res, 9999)
		}
	}

	for _, zero := range zeroList {
		for i := range s {
			//fmt.Println(zero - i)
			if int(math.Abs(float64(zero-i))) < res[i] {
				res[i] = int(math.Abs(float64(zero-i)))
			}
		}
	}
	return res
}

func main() {
	s := "loveleetcode"
	c := byte('e')
	fmt.Println(shortestToChar(s, c))
}
