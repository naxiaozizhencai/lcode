package main

import "fmt"

func main() {
	nums := []int{0,0,1}
	moveZeroes(nums)

	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	var zeroNums int
	for i := 0; i < len(nums); {

		if nums[i] == 0 {
			zeroNums++
			nums = append(nums[0:i], nums[i+1:]...)

		} else {
			i++
		}
	}


	for {
		if zeroNums == 0 {
			break
		}
		nums = append(nums, 0)
		zeroNums--
	}

}
