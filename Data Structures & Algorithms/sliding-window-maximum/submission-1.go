func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int,0)
	dequeue := make([]int,0)

	for i := range nums{
		if len(dequeue) > 0 && dequeue[0] <= i-k {
			dequeue = dequeue[1:]
		}

		for len(dequeue) > 0 && nums[dequeue[len(dequeue)-1]] < nums[i]{
			dequeue = dequeue[:len(dequeue)-1]
		}

		dequeue = append(dequeue, i)

		if i >= k-1{
			res = append(res, nums[dequeue[0]])
		}

	}

	return res
}
