func isValidSudoku(board [][]byte) bool {
    var rows, cols, subBoxes [9][9]bool

    for row := 0; row < 9; row++{
        for col := 0; col < 9; col++{
            if board[row][col] == '.'{
                continue
            }

            idx := board[row][col] - '1'

            subBoxIdx := (row / 3) * 3 + (col / 3)

            if subBoxes[subBoxIdx][idx] || rows[row][idx] || cols[col][idx]{
                return false
            }

            subBoxes[subBoxIdx][idx] = true
            rows[row][idx] = true
            cols[col][idx] = true
        }
    }

    return true
}