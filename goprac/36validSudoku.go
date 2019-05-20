func isValidSudoku(board [][]byte) bool {
    //fmt.Println("CheckEachThreeByThree: ",checkEachThreeByThree(board))
    //fmt.Println("checkLines: ", checkLines(board))
    if !checkEachThreeByThree(board) || !checkLines(board) {
        return false
    }
    return true
}

func checkThreeByThreeQuadrant(row_min, row_max, col_min, col_max int, board [][]byte) bool {
   // fmt.Println(row_min, row_max, col_min, col_max)
    boolArr := map[byte]bool{}
    for i := row_min; i < row_max; i++ {
        for j := col_min; j < col_max; j++ {
            if board[i][j] != '.' {
                if boolArr[board[i][j]] == false {
                    boolArr[board[i][j]] = true
                } else {
                    return false
                } 
            }
        }
    }
    return true
}

func checkEachThreeByThree (board [][]byte) bool {
    
    for i := 0; i <= 6; i+=3 {
        for j := 0; j <= 6; j+=3 {
            if !checkThreeByThreeQuadrant(i,i+3,j,j+3,board){
                return false
            }
        }
    }
    return true
    
}

func checkRow(row int, board [][]byte) bool {
    boolArr := map[byte]bool{}
    for i := 0; i < 9; i++ {
        if board[row][i] != '.' {
            if boolArr[board[row][i]] == false {
                boolArr[board[row][i]] = true
            } else {
                return false
            }
        }
    }
    return true
}

func checkCol(col int, board [][]byte) bool {
     boolArr := map[byte]bool{}
    for i := 0; i < 9; i++ {
        if board[i][col] != '.' {
            if boolArr[board[i][col]] == false {
                boolArr[board[i][col]] = true
            } else {
                return false
            }
        }
    }
    return true
}

func checkLines(board [][]byte) bool {
     for i := 0; i < 9; i++ {
         if !checkRow(i, board) {
             return false
         }   
         if !checkCol(i, board) {
             return false
         }   
     }
    return true
}
