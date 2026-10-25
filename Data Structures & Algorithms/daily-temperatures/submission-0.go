func dailyTemperatures(temperatures []int) []int {
	stack := make([]int,0)
	result := make([]int,len(temperatures))

	for i := len(temperatures)-1; i >= 0; i--{
		for len(stack) > 0 {
			stackTop := stack[len(stack)-1]
			stackTopTemp:= temperatures[stackTop]
			
			if stackTopTemp > temperatures[i]{
				result[i] = stackTop - i
				break
			}

			stack = stack[:len(stack)-1]
		}

		stack = append(stack,i)
	}

	return result
}
