func longestConsecutive(nums []int) int {
	hashMap := make(map[int]struct{})
	for _,num := range nums{
		hashMap[num] = struct{}{}
	}

	maxConsecutive := 0
	for key, _ := range hashMap{
		consecutive := 1
		left, right := key, key
		delete(hashMap, key)

		for{
			right++
			if _,exists := hashMap[right]; exists{
				delete(hashMap, right)
				consecutive++
			}else{
				break
			}
		}

		for{
			left--
			if _ , exists := hashMap[left]; exists{
				delete(hashMap, left)
				consecutive++
			}else{
				break
			}
		}

		if consecutive > maxConsecutive{
			maxConsecutive = consecutive
		}
	}

	return maxConsecutive
}
