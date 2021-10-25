package main

import "fmt"

func game(guess []int, answer []int) int {
	count := 0
	for i, v := range guess {
		if v == answer[i] {
			count++
		}
	}
	return count
}

func main() {
	guess := []int{1,2,3}
	answer := []int{1,2,3}

	fmt.Println(game(guess, answer))

}
