package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func getKthMagicNumber(k int) int {
	//numList := [k+1]int{1}  // array bound must be a constant expression code
	numList := make([]int, k+1)
	p3, p5, p7 := 0, 0, 0
	numList[0] = 1
	for i := 1; i < k; i++ {
		numList[i] = min(min(numList[p3]*3, numList[p5]*5), numList[p7]*7)
		if numList[i] == numList[p3]*3 {
			p3++
		}
		if numList[i] == numList[p5]*5 {
			p5++
		}
		if numList[i] == numList[p7]*7 {
			p7++
		}
	}
	return numList[k-1]
}

func main() {
	k := 5
	fmt.Println(getKthMagicNumber(k))
}
