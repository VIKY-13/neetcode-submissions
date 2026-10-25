func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _,n := range nums{
		freqMap[n]++
	}

	freqArr := make([][]int,len(nums)+1)
	for n, freq := range freqMap{
		freqArr[freq] = append(freqArr[freq], n)
	}

	res := make([]int,0)
	for row:=len(freqArr)-1; row>=0; row--{
		for n := len(freqArr[row])-1; n >=0; n--{
			res = append(res, freqArr[row][n])
			
			if len(res) == k {
                return res
            }
		}
	}

	return res
}
