func searchMatrix(matrix [][]int, target int) bool {
	low, high := 0, (len(matrix)* len(matrix[0]))-1
	n := len(matrix[0])

	for low <= high{
		mid := low + ((high-low)/2)
		row := mid / n
		col := mid % n

		if target == matrix[row][col]{
			return true
		}

		if target < matrix[row][col]{
			high = mid-1
		}else{
			low = mid+1
		}
	}

	return false
}
