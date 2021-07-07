package main

import "fmt"

func productExceptSelf(nums []int) []int {

	result := make([]int, len(nums))

	product := 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = product
		product *= nums[i]
	}
	product = 1
	for i := 0; i <= len(nums)-1; i++ {
		result[i] *= product
		product *= nums[i]
	}
	return result
}

func main() {
	input := []int{2, 4, 6, 8}
	fmt.Println(productExceptSelf(input))
}
