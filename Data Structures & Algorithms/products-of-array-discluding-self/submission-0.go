func productExceptSelf(nums []int) []int {
	res := make([]int,len(nums))

	res[0] = 1
	for i := 1; i < len(nums); i++{
		res[i] = res[i-1] * nums[i-1]
	}
	// [5,3,2]
	// [1,5,15]



	postFix := 1
	for i := len(res)-1; i >=0; i--{
		res[i] *= postFix
		postFix *= nums[i]
	}

	return res
}
