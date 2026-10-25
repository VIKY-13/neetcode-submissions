func maxArea(heights []int) int {
	i, j := 0, len(heights)-1
	maxArea := 0

	for i < j {
		width := j-i
		area := min(heights[i],heights[j]) * width

		if area > maxArea{
			maxArea = area
		}

		if heights[i] < heights[j]{
			i++
		}else{
			j--
		}
	}

	return maxArea
}

func min(a,b int)int{
	if a < b{
		return a
	}
	return b
}