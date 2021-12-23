package main

import "fmt"

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	//words := []string{"abc","car","ada","racecar","cool"}
	words := []string{"notapalindrome","racecar"}
	//words := []string{"def", "ghi"}
	fmt.Println(firstPalindrome(words))
}
