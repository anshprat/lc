func search(nums []int, target int) int {
	numMap := make(map[int]int)
	for k, v := range nums {
		numMap[v] = k
	}
	if _, ok := numMap[target]; ok {
		return numMap[target]
	}
	return -1
}
