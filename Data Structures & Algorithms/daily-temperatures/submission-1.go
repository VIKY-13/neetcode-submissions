func dailyTemperatures(temperatures []int) []int {
	stack := make([]int,0)
	result := make([]int,len(temperatures))

	for i, temp := range temperatures{
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]]{
			prevIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIdx] = i - prevIdx
		}
		stack = append(stack, i)
	}
	return result
}
