package solution121

// ============================================================================
// 121. Best Time to Buy and Sell Stock
// URL: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// ============================================================================

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 1 {
		return 0
	}
	earn := prices[0]
	m := 0
	for i := 1; i < size; i++ {
		if earn > prices[i] {
			earn = prices[i]
		} else if prices[i]-earn > m {
			m = prices[i] - earn
		}
	}
	return m
}
