func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	result := make([]int,2)

	for left < right{
		sum := numbers[left] + numbers[right]
		if sum == target{
			result[0], result[1] = left+1, right+1
			break
		}

		if sum > target{
			right--
		}else{
			left++
		}
	}

	return result
}
