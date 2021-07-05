package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	var (
		maxProfit int
		minPrice  = prices[0]
	)

	for _, price := range prices {
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}

	return maxProfit
}

func main() {
	input := []int{99, 7, 100, 2, 10}
	fmt.Println(maxProfit(input))
}
