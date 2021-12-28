package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	res := []int{}
	// 如果是数字，直接返回
	tmp, err := strconv.Atoi(expression)
	if err == nil {
		return []int{tmp}
	}

	for i, v := range expression {
		tmpC := string(v)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
			// 如果是运算符，则计算左右两边的算是
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int
					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					res = append(res, addNum)
				}
			}
		}
	}
	return res
}

func main() {
	//expression := "2-1-1"
	expression := "2*3-4*5"
	fmt.Println(diffWaysToCompute(expression))
}
