package main

import "fmt"

// 暴力破解超内存了
//func maxProfit(prices []int) int {
//	profitList := []int{}
//	for i := 0; i < len(prices) - 1; i++ {
//		for j := i+1; j < len(prices); j++ {
//			if prices[i] < prices[j] {
//				profitList = append(profitList, prices[j] - prices[i])
//			}
//		}
//	}
//
//	profit := 0
//	for _, v := range profitList {
//		if v > profit {
//			profit = v
//		}
//	}
//	return profit
//}

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if maxProfit < price - minPrice {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func main() {
	//prices := []int{7,1,5,3,6,4}
	prices := []int{2,4,1}
	fmt.Println(maxProfit(prices))
}
