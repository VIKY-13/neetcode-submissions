func subsets(nums []int) [][]int {
	res := make([][]int,0)
	subset := make([]int,0)
	var backtrack func(i int)

	backtrack = func(i int){
		if i == len(nums){
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		subset = append(subset, nums[i])
		backtrack(i+1)

		subset = subset[:len(subset)-1]
		backtrack(i+1)
	}

	backtrack(0)
	return res
}
