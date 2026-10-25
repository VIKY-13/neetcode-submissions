func findDuplicate(nums []int) int {
    for i := 0; i < len(nums); i++{
		val := nums[i]

		if val < 0{
			val *= -1
		}

		if nums[val] < 0{
			return val
		}

		nums[val] *= -1
	}

	return 0
}
