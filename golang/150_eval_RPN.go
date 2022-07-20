package main

import (
	"fmt"
	"strconv"
)

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func evalRPN(tokens []string) int {
	var stack []string
	for i := 0; i < len(tokens); i++ {
		if isNum(tokens[i]) {
			stack = append(stack, tokens[i])
		} else {
			tmp := 0
			v1, _ := strconv.Atoi(stack[len(stack)-2])
			v2, _ := strconv.Atoi(stack[len(stack)-1])
			if tokens[i] == "+" {
				tmp =  v1 + v2
			} else if tokens[i] == "-" {
				tmp = v1 - v2
			} else if tokens[i] == "*" {
				tmp = v1 * v2
			} else if tokens[i] == "/" {
				tmp = v1 / v2
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(tmp))
		}
		//fmt.Println(stack)
	}
	res, _ := strconv.Atoi(stack[0])
	return res
}

func main() {
	//tokens := []string{"2","1","+","3","*"}  // (2+1) * 3 = 9
	tokens := []string{"4","13","5","/","+"}  // (2+1) * 3 = 9
	fmt.Println(evalRPN(tokens))
}