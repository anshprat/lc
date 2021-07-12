package main

import (
	"fmt"
	"math/bits"
)

func getSum(a int, b int) int {
	res, _ := bits.Add(uint(a), uint(b), 0)
	return int(res)
}

func main() {
	fmt.Println(getSum(1, 2))
}
