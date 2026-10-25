func evalRPN(tokens []string) int {
	stack := make([]int, 0 ,len(tokens))

	for i := 0; i < len(tokens); i++{
		if tokens[i] != "+" && tokens[i] != "-" && 
			tokens[i] != "*" && tokens[i] != "/" {
				val, _ := strconv.Atoi(tokens[i])
				stack = append(stack, val)
				continue
			}
		
		left := stack[len(stack)-2]
		right := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		switch tokens[i]{
			case "+":
				stack = append(stack, left + right)
			case "-":
				stack = append(stack, left - right)
			case "*":
				stack = append(stack, left * right)
			case "/":
				stack = append(stack, left / right)
		}
	}

	return stack[0]
}
