package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	l := 0
	r := len(height) - 1

	var area int
	var temp_area int

	for l < r {

		if height[l] > height[r] {
			temp_area = height[r] * (r - l)
			r--
		} else {
			temp_area = height[l] * (r - l)
			l++
		}
		fmt.Println(temp_area)
		if temp_area > area {
			area = temp_area
		}

	}

	return area
}
