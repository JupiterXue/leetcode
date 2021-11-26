package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "#" {
			if i == 0 {
				s = s[i+1:]
				i--
			} else if i == len(s) {
				s = s[:i-1]
				i--
			} else {
				s = s[:i-1] + s[i+1:]
				i--
				i--
			}
		}
	}

	for i := 0; i < len(t); i++ {
		if string(t[i]) == "#" {
			if i == 0 {
				t = t[i+1:]
				i--
			} else if i == len(t) {
				t = t[:i-1]
				i--
			} else {
				t = t[:i-1] + t[i+1:]
				i--
				i--
			}
		}
	}

	fmt.Println(s)
	fmt.Println(t)
	return s == t
}

func main() {
	s := "a##c"
	t := "a##c"
	fmt.Println(backspaceCompare(s, t))
}
