type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// 5#Hello5World0
	encodedStr := ""
	for _, str := range strs{
		encoded := fmt.Sprintf("%d#%s",len(str),str)
		encodedStr += encoded
	}

	return encodedStr
}

func (s *Solution) Decode(encoded string) []string {
	decodedVal := make([]string,0)
	
	i:=0
	for i < len(encoded){
		j := i
		for encoded[j] != '#'{
			j++
		}

		length,_ := strconv.Atoi(encoded[i:j])

		start := j+1
		end := start+length

		decodedVal = append(decodedVal, encoded[start:end])

		i = end
	}

	return decodedVal
}
