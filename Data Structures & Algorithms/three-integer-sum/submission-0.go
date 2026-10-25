// import "sort"
func threeSum(nums []int) [][]int {
	res := make([][]int,0)

	sort.Ints(nums)

	for i, num := range nums{
		if num > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1]{
			continue
		}

		left := i+1
		right := len(nums)-1
		for left < right{
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0{
				triplet := []int{nums[i], nums[left], nums[right]}
				res = append(res, triplet)

				for left < right && nums[left] == nums[left+1]{
					left++
				}

				for left < right && nums[right] == nums[right-1]{
					right--
				}

				left++
				right--
			} else if sum > 0{
				right--
			}else{
				left++
			}
		}
	}

	return res
}
