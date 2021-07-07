package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	var max int = -1000000
	for i := range nums {
		if i > 0 && nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max

}

func main() {
	input := []int{5, -1, -1, 9}
	fmt.Println(maxSubArray(input))
}
