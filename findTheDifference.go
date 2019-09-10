package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}

func findTheDifference(s string, t string) byte {
	var result uint8

	for _, v := range []rune(s) {
		result -= uint8(v)
	}

	for _, v := range []rune(t) {
		result += uint8(v)
	}

	return result
}
