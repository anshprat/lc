package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxLen := len(height) - 1
	maxArea := 0
	for i, j := 0, maxLen; i < j; {
		fmt.Println(i, j)
		if height[i] < height[j] {
			maxArea = max(maxArea, height[i]*(j-i))
			i++
		} else {
			maxArea = max(maxArea, height[j]*(j-i))
			j--
		}
	}
	return maxArea
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	input := []int{1, 3, 2, 5, 25, 24, 5}
	fmt.Println(maxArea(input))
}
