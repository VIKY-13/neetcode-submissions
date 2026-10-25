func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int,0)
	deque := make([]int,0)

	for i := range nums{
		if len(deque) > 0 && deque[0] <= i-k{
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i]{
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i-k+1 >= 0{
			res = append(res, nums[deque[0]])
		}
	}

	return res
}
