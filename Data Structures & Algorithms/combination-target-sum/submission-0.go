func combinationSum(nums []int, target int) [][]int {
    res := make([][]int,0)
	path := make([]int,0)

	var backtrack func(idx int)

	backtrack = func(idx int){
		if idx == len(nums){
			return
		}

		sum := sum(path)
		if sum == target{
			// fmt.Println("sum - ", path, sum)
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		if sum > target{
			return
		}

		// fmt.Println("before internal backt - ",path)
		path = append(path, nums[idx])
		backtrack(idx)

		path = path[:len(path)-1]
		// fmt.Println("after internal backt - ",path)

		backtrack(idx+1)
	}

	backtrack(0)
	return res
}

func sum(arr []int) int{
	sum := 0
	for _, val := range arr{
		sum += val
	}

	return sum
}
