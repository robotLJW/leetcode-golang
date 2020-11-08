package _22_best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	disparity := make([]int, len(prices)-1)
	for i := 0; i < len(disparity); i++ {
		disparity[i] = prices[i+1] - prices[i]
	}
	profit := 0
	for i := 0; i < len(disparity); i++ {
		if disparity[i] > 0 {
			profit = profit + disparity[i]
		}
	}
	return profit
}
