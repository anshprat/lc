package main

import (
	"fmt"
	"math/bits"
)

func countBits(n int) []int {
	res := []int{}
	for i := 0; i <= n; i++ {
		res = append(res, bits.OnesCount(uint(i)))
	}
	return res

}

func main() {
	fmt.Println(countBits(5))
}
