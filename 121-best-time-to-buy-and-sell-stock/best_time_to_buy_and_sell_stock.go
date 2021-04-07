package _21_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit, minValue := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		} else if prices[i]-minValue > maxProfit {
			maxProfit = prices[i] - minValue
		}
	}
	return maxProfit
}
