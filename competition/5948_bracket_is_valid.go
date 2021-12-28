package main

import "fmt"

func canBeValid(s string, locked string) bool {

}

func main() {
	s := "))()))"
	locked := "010100"
	fmt.Println(canBeValid(s, locked))
}