package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] ++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}

	}

	result := make([]int , 1)
	result[0] = 1
	result = append(result, digits...)

	return result

}
