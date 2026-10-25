func largestRectangleArea(heights []int) int {
	stack := make([]int, 0 , len(heights))
	heights = append(heights, 0)
	maxArea := 0

	for i := 0; i< len(heights); i++{
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]]{
			top := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			left := -1
			if len(stack) > 0{
				left = stack[len(stack)-1]
			}

			width := i-left-1
			area := top * width

			if area > maxArea{
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}
