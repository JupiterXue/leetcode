package main

import "fmt"

func maxProfit2(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

// AC!!!
//func maxProfit2(prices []int) int {
//	allProfit := 0
//	maxProfit := 0
//	buyPrice := prices[0]
//	buyState := false
//	// 先找最小的价格，找到了买入
//	// 再找最大的利润，找到后买入下一个最小的
//	// 再找最大的利润，找到后。。。重复
//	for i, price := range prices {
//		if i == 0 {
//			continue
//		}
//		if price < buyPrice && buyState == false {
//			buyPrice = price
//			buyState = true
//		} else {
//			if maxProfit < price - buyPrice { // 如果当前利润更大，记录
//				maxProfit = price - buyPrice
//				buyState = true
//			} else { // 这里是下一次价格，说明 maxProfit 达到了最大，累计并清零
//				allProfit += maxProfit
//				maxProfit = 0
//				buyPrice = prices[i]
//				buyState = true
//			}
//		}
//		fmt.Print("当前价：")
//		fmt.Print(price)
//		fmt.Print("，买入价：")
//		fmt.Print(buyPrice)
//		fmt.Print("，利润：")
//		fmt.Print(maxProfit)
//		fmt.Print("，总计利润：")
//		fmt.Print(allProfit)
//		fmt.Println()
//	}
//	return allProfit+maxProfit
//}

func main() {
	//prices := []int{7,1,5,3,6,4}
	//prices := []int{1,2,3,4,5}
	prices := []int{1,2,4,2,5,7,2,4,9,0}
	//prices := []int{2,1,2,0,1}
	//prices := []int{2,4,1,7}
	fmt.Println(maxProfit2(prices))
}
