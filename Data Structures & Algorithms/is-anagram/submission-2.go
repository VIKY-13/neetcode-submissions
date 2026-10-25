func isAnagram(s string, t string) bool {
	letterToCountMap := map[rune]int{}

	for _,str := range s{
		letterToCountMap[str] += 1
	}

	for _, str := range t{
		count , exists := letterToCountMap[str]
		if !exists {
			return false
		}
		// fmt.Println(count)
		
		count -= 1
		letterToCountMap[str] = count

		if count <= 0{
			delete(letterToCountMap,str)
		}
	}

	if len(letterToCountMap) > 0{
		return false
	}

	return true
}
