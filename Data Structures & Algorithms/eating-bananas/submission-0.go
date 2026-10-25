func minEatingSpeed(piles []int, h int) int {
	maxPile := piles[0]
	for _, p := range piles{
		if p > maxPile{
			maxPile = p
		}
	}
	low, high := 1,maxPile
	ans := maxPile

	for low <= high{
		mid := low + (high-low) / 2

		hours := 0
		for _,p := range piles{
			hours += (p+mid-1)/mid
		}

		if hours <= h{
			ans = mid
			high = mid-1
		}else {
			low = mid+1
		}

	}

	return ans
}
