type Solver struct {
    seenRow [9][9]bool
    seenCol [9][9]bool
    seenBox [3][3][9]bool
    board [9][9]int
}

func (s *Solver) save(row, col, num int) {
    s.board[row][col] = num
    s.seenRow[row][num-1] = true
    s.seenCol[col][num-1] = true
    s.seenBox[row/3][col/3][num-1] = true
}

func (s *Solver) invalid(row, col, num int) bool {
    cellNotEmpty := s.board[row][col] != -1
    inRow := s.seenRow[row][num-1]
    inCol := s.seenCol[col][num-1]
    inBox := s.seenBox[row/3][col/3][num-1]
    return cellNotEmpty || inRow || inCol || inBox
}

func (s *Solver) findEmpty() (int, int, bool) {
    for i := range s.board {
        for j, val := range s.board[i] {
            if val == -1 {
                return i, j, true
            }
        }
    }
    return -1, -1, false
}

func (s *Solver) reset(row, col, num int) {
    s.board[row][col] = -1
    s.seenRow[row][num-1] = false
    s.seenCol[col][num-1] = false
    s.seenBox[row/3][col/3][num-1] = false
}

func (s *Solver) backtrack() bool {
    row, col, emptyFound := s.findEmpty()
    if !emptyFound {
        return true
    }

    for num := 1; num <= 9; num++ {
        if s.invalid(row, col, num) {
            continue
        }

        s.save(row, col, num)
        
        if s.backtrack() {
            return true
        }

        s.reset(row, col, num)
    }

    return false
}

func solveSudoku(board [][]byte)  {
    s := Solver{}

    for row := range board {
        for col, byteVal := range board[row] {
            if byteVal == '.' {
                s.board[row][col] = -1
                continue
            }
            s.save(row, col, int(byteVal-'0'))
        }
    }

    if done := s.backtrack(); !done {
        return
    }

    for row := range s.board {
        for col, intVal := range s.board[row] {
            board[row][col] = byte(intVal)+'0'
        }
    }
}
