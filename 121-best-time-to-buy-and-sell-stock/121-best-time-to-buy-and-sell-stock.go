/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sat Sep 10 2022
 */

// REFERENCE: https://leetcode.com/problems/best-time-to-buy-and-sell-stock

package main

func maxProfit(prices []int) int {
	maxProfitVal := 0
	lowestBuyVal := prices[0]

	for _, price := range prices {
		profit := price - lowestBuyVal
		if profit > maxProfitVal {
			maxProfitVal = profit
		}
		if price < lowestBuyVal {
      lowestBuyVal = price;
    }
	}
	return maxProfitVal
}

// func main() {
// 	arr := [5]int{1, 5, 3, 6, 4}
// 	maxProfit(arr[:])
// }
