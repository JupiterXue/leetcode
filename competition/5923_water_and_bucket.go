package main

import (
	"fmt"
	"strings"
)

func minimumBuckets(street string) int {
	if strings.Contains(street, "HHH") {
		return -1
	}
	if len(street) > 1 {
		if street[:2] == "HH" || street[len(street)-2:] == "HH" {
			return -1
		}
	}
	newStreet := strings.Replace(street, "H.H", "", -1)
	count := (len(street) - len(newStreet)) / 3
	if len(newStreet) == 0 {
		return count
	}
	if strings.Contains(newStreet, ".") {
		for _, v := range newStreet {
			char := string(v)
			if char == "H" {
				count++
			}
		}
	} else {
		return -1
	}
	return count
}

func main() {
	//street := "H..H"
	//street := ".H.H."
	//street := ".HHH."
	//street := "H"
	//street := "HH......"
	//street := "...HH"
	//street := "...HH"
	street := "H.H"
	fmt.Println(minimumBuckets(street))
}