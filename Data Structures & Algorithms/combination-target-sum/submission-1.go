func combinationSum(nums []int, target int) [][]int {
    res := make([][]int,0)
	path := make([]int,0)

	var backtrack func(idx, currSum int)

	backtrack = func(idx, currSum int){
		if idx == len(nums) || currSum > target{
			return
		}

		if currSum == target{
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		path = append(path, nums[idx])
		backtrack(idx, currSum + nums[idx])

		path = path[:len(path)-1]

		backtrack(idx+1, currSum)
	}

	backtrack(0,0)
	return res
}
