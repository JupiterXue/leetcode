package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	for i := 1; i < len(arr) - 1; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}

func main() {
	arr := []int{0,2,1,0}
	fmt.Println(peakIndexInMountainArray(arr))
}
