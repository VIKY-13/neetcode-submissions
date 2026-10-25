func evalRPN(tokens []string) int {
	stack := make([]int,0)

	for _,token := range tokens{
		if token != "+" && token != "-" &&
			token != "*" && token != "/"{
				val, _ := strconv.Atoi(token)
				stack = append(stack, val)
				continue
		}

		top := stack[len(stack)-1]
		prevTop := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch token{
			case "+":
				stack = append(stack, prevTop+top)
				// fmt.Println("in add - ",stack, top, prevTop)
			case "-":
				stack = append(stack, prevTop-top)
			case "*":
				stack = append(stack, prevTop*top)
				// fmt.Println("in mul",stack, top, prevTop)
			case "/":
				stack = append(stack, prevTop/top)
                // fmt.Println("in div - ",stack, top, prevTop)
		}
	}

	return stack[0]
}
