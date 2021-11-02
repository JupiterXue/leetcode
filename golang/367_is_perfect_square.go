package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(num int) bool {
	square := math.Sqrt(float64(num))
	square_int := int( math.Sqrt(float64(num)))
	if float64(square_int) == square {
		return true
	} else {
		return false
	}
}

func main() {
	num := 14
	fmt.Println(isPerfectSquare(num))
}
