package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1,2,3,4,5,6,9}))
}
func maxProfit(prices []int) int {

	totalProfit := 0

	if len(prices) <= 1{
		return totalProfit
	}

	for i:=1 ;i<len(prices); i++{
		if prices[i] > prices[i - 1] {
			totalProfit += prices[i] - prices[i - 1];
		}
	}

	return totalProfit
}
