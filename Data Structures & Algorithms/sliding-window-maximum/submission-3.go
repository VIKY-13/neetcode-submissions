func maxSlidingWindow(nums []int, k int) []int {
    res, deq := make([]int,0), make([]int,0)

	for i := range nums{
		if len(deq) > 0 && deq[0] <= i-k{
			deq = deq[1:]
		}

		for len(deq) > 0 && nums[deq[len(deq)-1]] < nums[i]{
			deq = deq[:len(deq)-1]
		}

		deq = append(deq, i)

		if i >= k-1{
			res = append(res, nums[deq[0]])
		}
	}

	return res
}
