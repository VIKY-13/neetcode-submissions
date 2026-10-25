func characterReplacement(s string, k int) int {
	left := 0
	maxCount, maxUnique := 0, 0
	count := make(map[byte]int)

	for right := 0; right < len(s); right++{
		count[s[right]]++

		if count[s[right]] > maxCount{
			maxCount = count[s[right]]
		}

		for (right-left+1) - maxCount > k{
			count[s[left]]--
			left++
		}

		if right-left+1 > maxUnique{
			maxUnique = right-left+1
		}
	}
	return maxUnique
}
