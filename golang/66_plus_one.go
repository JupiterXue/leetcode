package main

import "fmt"

func plusOne(digits []int) []int {
	digits[len(digits)-1] = digits[len(digits)-1] + 1
	i := len(digits) - 1
	current := digits[i]
	count := 0 // 进位
	for current  == 10 {
		digits[i] = 0
		count++
		i--
		if i < 0 {
			break
		}
		current = digits[i] + count
		count--
		if current != 10 {
			digits[i]++
		}
	}
	if count > 0 {
		tmp := []int{1}
		digits = append(tmp, digits...)
	}
	return digits
}

func main() {
	digits := []int{9,8,9}
	fmt.Println(plusOne(digits))
}
