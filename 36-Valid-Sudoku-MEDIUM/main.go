func isValidSudoku(board [][]byte) bool {
    seenRow := [9][9]bool{}
    seenCol := [9][9]bool{}
    seenBox := [3][3][9]bool{}

    for row := range board {
        for col := range board[row] {
            num, err := strconv.Atoi(string(board[row][col]))
            if err != nil {
                continue
            }

            if seenRow[row][num-1] || seenCol[col][num-1] || seenBox[row/3][col/3][num-1] {
                return false
            }
            seenRow[row][num-1] = true
            seenCol[col][num-1] = true
            seenBox[row/3][col/3][num-1] = true
        }
    }
    return true
}
