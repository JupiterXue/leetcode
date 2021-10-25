package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32 && num > 0; i++ {
		if  num & 1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

func main() {
	num := uint32(00000000000000000000000010000000)
	fmt.Println(hammingWeight(num))
}
