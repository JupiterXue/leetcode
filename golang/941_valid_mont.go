package main

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	maxValue := 0
	maxIndex := 0
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
		if arr[i] > maxValue {
			maxValue = arr[i]
			maxIndex = i
		}
	}

	for i := 0; i < len(arr) - 1; i++ {
		if i < maxIndex {
			if arr[i] >= arr[i+1] {
				return false
			}
		} else {
			if arr[i] <= arr[i+1] {
				return false
			}
		}

		if arr[i] > maxValue {
			maxValue = arr[i]
			maxIndex = i
		}
	}

	if arr[maxIndex] > arr[len(arr) -1] && maxIndex > 0  {
		return true
	} else {
		return false
	}
}
//
//func main() {
//	arr := []int{9,8,7,6,5,4,3,2,1,0}
//	fmt.Println(validMountainArray(validMountainArray(arr)))
//}
