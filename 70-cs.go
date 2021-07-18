package main

import "fmt"

func climbStairs(n int) int {
	n0, n1, n2 := 1, 1, 1
	for i := 0; i < n; i++ {
		n0 = n1
		n1 = n1 + n2
		n2 = n0
		fmt.Println(n, i, n0, n1, n2)
	}
	return n2
}

func main() {
	fmt.Println(climbStairs(0))
}
