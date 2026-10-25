func lengthOfLongestSubstring(s string) int {
	charToIdxMap := make(map[rune]int)
	left := 0
	subStr := 0

	for i, char := range s{
		if idx, exists := charToIdxMap[char]; exists && left <= idx {
			left = idx+1
		}
		charToIdxMap[char] = i

		currSubStr := i - left + 1
		if currSubStr > subStr{
			subStr = currSubStr
		}
	}

	return subStr
}
