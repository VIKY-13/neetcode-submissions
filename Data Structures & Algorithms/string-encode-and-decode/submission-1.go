type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	endcodedStr := ""

	for _, str := range strs{
		fmtStr := fmt.Sprintf("%d#%s", len(str), str)
		endcodedStr += fmtStr
	}

	return endcodedStr
}

func (s *Solution) Decode(encoded string) []string {
	decoded := make([]string,0)
	i := 0

	for i < len(encoded){
		j := i
		for encoded[j] != '#'{
			j++
		}

		strSpan, _ := strconv.Atoi(encoded[i:j])

		start := j+1
		end := start+strSpan
		decoded = append(decoded, encoded[start:end])

		i = end
	}

	return decoded
}
