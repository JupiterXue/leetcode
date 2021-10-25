package main

func moveZeroes(nums []int)  {
	tryTime := len(nums)
	for i := 0; i < len(nums); i++ {
		if tryTime <= 0 {
			break
		}
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			i--
		}
		tryTime--
	}
}

func main() {
	//nums := []int{0,1,0,3,12}
	nums := []int{0,0,1}
	moveZeroes(nums)
	//fmt.Println(nums)
}
