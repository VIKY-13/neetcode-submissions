func minWindow(s string, t string) string {
    if len(s) < len(t){
		return ""
	}

	tCount := make(map[byte]int)
	window := make(map[byte]int)

	for i := range t{
		tCount[t[i]]++
	}

	need := len(tCount)
	left, have := 0, 0
	res := [2]int{}
	resLen := len(s)+1

	for right := 0; right < len(s); right++{
		window[s[right]]++

		if count, ok := tCount[s[right]]; ok && count == window[s[right]]{
			have++
		}

		for have == need{
			winLen := right - left + 1
			if winLen < resLen{
				resLen = winLen
				res[0], res[1] = left, right
			}

			window[s[left]]--
			if count, ok := tCount[s[left]]; ok && window[s[left]] < count {
				have--
			}
			left++
		}
	}

	if resLen > len(s){
		return ""
	}

	return s[res[0]:res[1]+1]
}
