func maxProfit(prices []int) int {
	left,right := 0,0
	// minPrice := prices[left]
	maxPrice := 0
	for right < len(prices){
		if prices[right] - prices[left] > maxPrice{
			maxPrice = prices[right] - prices[left]
		}

		if prices[right] < prices[left]{
			left = right
		}

		right++
	}

	return maxPrice
}
