package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 23, 45}, 68))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		tempValue := target - nums[i]
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if tempValue == nums[j] {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
