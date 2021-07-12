package main

import (
	"fmt"
	"math/bits"
)

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}

func main() {
	nums := "00000010100101000001111010011100"
	fmt.Println(reverseBits(uint32(nums)))
}
