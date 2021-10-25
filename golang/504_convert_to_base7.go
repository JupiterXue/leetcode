package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	res := strconv.FormatInt(int64(num), 7)
	return string(res)
}

func main() {
	num := 100
	fmt.Println(convertToBase7(num))
}
