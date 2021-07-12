package main

import (
	"fmt"
	"math/bits"
)

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)

}

func main() {
	fmt.Println(hammingWeight("11111111111111111111111111111101"))
}
