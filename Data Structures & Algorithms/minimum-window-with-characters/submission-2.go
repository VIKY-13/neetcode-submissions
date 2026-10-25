func minWindow(s string, t string) string {
    tMap := make(map[byte]int)
	for i := range t{
		tMap[t[i]]++
	}
	need := len(tMap)

	have := 0
	res := [2]int{}
	left := 0
	window := make(map[byte]int)
	minWindowLen := len(s)+1

	for right := 0; right < len(s); right++{
		window[s[right]]++
		if count, ok := tMap[s[right]]; ok && count == window[s[right]]{
			have++
		}

		for have == need{
			if right-left+1 < minWindowLen{
				res[0], res[1] = left,right
				minWindowLen = right-left+1
			}

			window[s[left]]--

			if count, ok := tMap[s[left]]; ok && window[s[left]] < count{
				have--
			}

			left++
		}
	}

	if minWindowLen == len(s)+1{
		return ""
	}

	return s[res[0]:res[1]+1]
}
