func solveNQueens(n int) [][]string {
    final := make([][]int, 0)
    board := make([]int, n)
    solve(n-1, board, &final)

    answer := make([][]string, len(final))

    for i, b := range final {
        answer[i] = make([]string, n)
        for j, chosenCol := range b {
            row := make([]byte, n)
            for k := range row {
                row[k] = '.'
            }
            row[chosenCol] = 'Q'
            answer[i][j] = string(row)
        }
    }

    return answer
}

func valid(i int, board []int) bool {
    for j := i + 1; j < len(board); j++ {
        if (
            board[j] == board[i] ||
            board[j] == board[i] + (j-i) ||
            board[j] == board[i] - (j-i)) {
            return false
        }
    }
    return true
}

func solve(i int, board []int, final *[][]int) {
    if i < 0 {
        *final = append(*final, append([]int(nil), board...))
        return
    }

    for j := range len(board) {
        board[i] = j
        if valid(i, board) {
            solve(i-1, board, final)
        }
    }
}
