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
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(input))
}
