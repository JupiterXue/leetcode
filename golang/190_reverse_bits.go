package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32 && num > 0; i++ {
		res |= num & num & 1 << (31 - i)
		num >>= 1
	}
	return res
}

func main() {
	num := uint32(00000010100101000001111010011100)
	//num := uint32(11111111111111111111111111111101)
	//fmt.Println(num)
	fmt.Println(reverseBits(num))
}
