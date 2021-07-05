package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var inputs = make(map[int]int)
	for c := 0; c < len(nums); c++ {
		if c > 0 {
			comp := target - nums[c]
			_, compExists := inputs[comp]
			if compExists {
				res := [2]int{inputs[comp], c}
				return (res[0:])
			}
		}
		inputs[nums[c]] = c
	}
	return (nums[0:])

}

func main() {
	nums := [2]int{3, 3}
	fmt.Println(twoSum(nums[0:], 6))
}
