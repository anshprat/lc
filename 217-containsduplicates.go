package main

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			return true
		}
		numMap[v] = v
	}
	return false
}
