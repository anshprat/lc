package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	//max(prices[len(prices)-1->1]-prices[len(prices)-2->0])
	//sp , cp, p
	// sd, cd
	// profit
	// var profit []int
	var maxProfit int = 0

	spMap := make(map[int]int)
	var dedupPrices []int
	for t := 0; t < len(prices); t++ {
		if t > 0 {
			if dedupPrices[len(dedupPrices)-1] != prices[t] {
				dedupPrices = append(dedupPrices, prices[t])
			}

		} else {
			dedupPrices = append(dedupPrices, prices[t])

		}
	}
	numDays := len(dedupPrices) - 1
	for sd := numDays; sd >= 1; sd = sd - 1 {
		for cd := sd - 1; cd >= 0; cd = cd - 1 {
			if dedupPrices[sd] > dedupPrices[cd] {
				sp := dedupPrices[sd] - dedupPrices[cd]
				_, spSeen := spMap[sp]
				if !spSeen {
					if dedupPrices[sd]-dedupPrices[cd] > maxProfit {
						maxProfit = int(dedupPrices[sd] - dedupPrices[cd])
					}
					spMap[sp] = sp
				}
			}
		}
	}

	return (maxProfit)
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(input))
}
