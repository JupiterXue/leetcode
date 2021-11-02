package main

import (
	"fmt"
	"math"
)

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

func main() {
	x := 2.1000
	n := 3
	fmt.Println(myPow(x, n))
}