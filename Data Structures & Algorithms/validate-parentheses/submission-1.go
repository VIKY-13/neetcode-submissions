func isValid(s string) bool {
	stack := make([]rune,0)
	bracketMap := map[rune]rune{
		'{':'}',
		'[':']',
		'(':')',
	}
    for _, char := range s{
		if char == '{' || char == '[' || char == '('{
			stack = append(stack, char)
			continue
		}
        if len(stack) == 0 {
            return false
        }
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if bracketMap[top] != char{
			return false
		}
	}

	if len(stack) > 0{
		return false
	}
	return true
}
