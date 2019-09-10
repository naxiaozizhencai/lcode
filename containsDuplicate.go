package main

import "fmt"

func main(){
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
}
func containsDuplicate(nums []int) bool {

	if len(nums) == 1 {
		return false
	}

	numsMap := make(map[int]int, len(nums))
	for i:=0;i<len(nums);i++{
		value, ok := numsMap[nums[i]]
		 if !ok {
			 numsMap[nums[i]] = 1
		 }else{
			 numsMap[nums[i]] += 1
			 if value + 1 >= 2 {
		 		return true
			}
		 }
	}
	return false
}
