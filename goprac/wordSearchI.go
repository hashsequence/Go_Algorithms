func exist(board [][]byte, word string) bool {
    w := []rune(word)
    dir := [4]Dir{Dir{0,1}, Dir{1,0}, Dir{0,-1}, Dir{-1,0}}
    
    for i, _ := range board {
        for j, _ := range board[i] {
            if rune(board[i][j]) == w[0] {
                board[i][j] = byte('#')
                if ifExists(&board, &w, &dir, 1, Dir{i,j}) {
                return true
                } 
                board[i][j] = byte(w[0])
                
            }
        }
    }
    return false
}

func ifExists(board *[][]byte, w *[]rune, dir *[4]Dir, height int, coord Dir) bool {
    if height == len(*w) {
        return true
    }
    for k, _ := range (*dir) {
        newCoord := Dir{coord.x + (*dir)[k].x, coord.y + (*dir)[k].y}
        if  newCoord.x < 0 || newCoord.y < 0 || newCoord.x >= len((*board)) || newCoord.y >= len((*board)[0]) {
            continue
        }
        if  rune((*board)[newCoord.x][newCoord.y]) == (*w)[height] {
             (*board)[newCoord.x][newCoord.y] = byte('#')
            if ifExists(board, w, dir, height+1, newCoord) {
              return true
            } 
            (*board)[newCoord.x][newCoord.y] = byte((*w)[height])
        }
    }
    return false
    
}


type Dir struct {
    x int
    y int
}

