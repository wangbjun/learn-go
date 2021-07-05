package main

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	minPrice := 7
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
