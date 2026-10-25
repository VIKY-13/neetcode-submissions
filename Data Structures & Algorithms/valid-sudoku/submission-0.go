func isValidSudoku(board [][]byte) bool {
	var rows, cols, subBox [9][9]bool

	for row := 0; row < 9; row++{
		for col := 0; col < 9; col++{
			if board[row][col] == '.'{
				continue
			}

			idx := board[row][col] - '1'
			val := (row/3) * 3 + (col / 3)

			if rows[row][idx] || cols[col][idx] || subBox[val][idx]{
				return false
			}

			rows[row][idx] = true
			cols[col][idx] = true
			subBox[val][idx] = true
		}
	}

	return true
}
