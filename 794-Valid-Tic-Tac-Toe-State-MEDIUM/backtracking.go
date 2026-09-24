func equal(a, b [][]byte) bool {
    for i := range a {
        if string(a[i]) != string(b[i]) {
            return false
        }
    }

    return true
}

func won(board [][]byte) bool {
    sameInPrimDiag, sameInSecDiag := true, true
    for i := range board {
        sameInCol, sameInRow := true, true

        for j := range board[i] {
            if board[i][j] == ' ' || board[i][j] != board[i][0] {
                sameInRow = false
            }
            if board[j][i] == ' ' || board[j][i] != board[0][i] {
                sameInCol = false
            }
        }

        if board[i][i] == ' ' || board[i][i] != board[0][0] {
            sameInPrimDiag = false
        }
        if board[i][2-i] == ' ' || board[i][2-i] != board[0][2] {
            sameInSecDiag = false
        }

        if sameInRow || sameInCol {
            return true
        }
    }
    return sameInPrimDiag || sameInSecDiag
}

func validTicTacToe(board []string) bool {
    var solve func() bool
    xTurn := true
    goalBoard := [][]byte{
        []byte(board[0]),
        []byte(board[1]),
        []byte(board[2]),
    }
    newBoard := [][]byte{
        []byte("   "),
        []byte("   "),
        []byte("   "),
    }

    solve = func() bool {
        if won(newBoard) {
            return equal(goalBoard, newBoard)
        }
        if equal(goalBoard, newBoard) {
            return true
        }
        for i := range newBoard {
            for j, sq := range newBoard[i] {
                if sq == ' ' {
                    if xTurn {
                        newBoard[i][j] = 'X'
                    } else {
                        newBoard[i][j] = 'O'
                    }
                    xTurn = !xTurn

                    if solve() {
                        return true
                    }

                    newBoard[i][j] = ' '
                    xTurn = !xTurn
                }
            }
        }
        return false
    }

    return solve()
}
