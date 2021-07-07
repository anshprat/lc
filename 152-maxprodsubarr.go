package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	maxPro := math.Inf(-1)
	curPro := float64(1)

	for i := range nums {
		curPro *= float64(nums[i])
		maxPro = math.Max(curPro, maxPro)
		if curPro == 0 {
			curPro = 1
		}
	}
	curPro = 1
	for i := len(nums) - 1; i > 0; i-- {
		curPro *= float64(nums[i])
		maxPro = math.Max(curPro, maxPro)
		if curPro == 0 {
			curPro = 1
		}
	}

	return int(maxPro)
}

func main() {
	// input := []int{-3, -1, -1}
	// input := []int{3, -1, 4}
	// input := []int{2, 3, -2, 4}
	input := []int{2, -5, -2, -4, 3}
	fmt.Println(maxProduct(input))
}
