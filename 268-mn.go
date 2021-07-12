package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)
	c := len(nums)
	if nums[c-1] == c-1 {
		return c
	}

	for k, v := range nums {
		if k != v {
			return k
		}
	}
	return 0
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
