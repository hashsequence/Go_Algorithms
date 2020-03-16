/*
909. Snakes and Ladders
Medium


On an N x N board, the numbers from 1 to N*N are written boustrophedonically starting from the bottom left of the board, and alternating direction each row.  For example, for a 6 x 6 board, the numbers are written as follows:


You start on square 1 of the board (which is always in the last row and first column).  Each move, starting from square x, consists of the following:

You choose a destination square S with number x+1, x+2, x+3, x+4, x+5, or x+6, provided this number is <= N*N.
(This choice simulates the result of a standard 6-sided die roll: ie., there are always at most 6 destinations, regardless of the size of the board.)
If S has a snake or ladder, you move to the destination of that snake or ladder.  Otherwise, you move to S.
A board square on row r and column c has a "snake or ladder" if board[r][c] != -1.  The destination of that snake or ladder is board[r][c].

Note that you only take a snake or ladder at most once per move: if the destination to a snake or ladder is the start of another snake or ladder, you do not continue moving.  (For example, if the board is `[[4,-1],[-1,3]]`, and on the first move your destination square is `2`, then you finish your first move at `3`, because you do not continue moving to `4`.)

Return the least number of moves required to reach square N*N.  If it is not possible, return -1.

Example 1:

Input: [
[-1,-1,-1,-1,-1,-1],
[-1,-1,-1,-1,-1,-1],
[-1,-1,-1,-1,-1,-1],
[-1,35,-1,-1,13,-1],
[-1,-1,-1,-1,-1,-1],
[-1,15,-1,-1,-1,-1]]
Output: 4
Explanation: 
At the beginning, you start at square 1 [at row 5, column 0].
You decide to move to square 2, and must take the ladder to square 15.
You then decide to move to square 17 (row 3, column 5), and must take the snake to square 13.
You then decide to move to square 14, and must take the ladder to square 35.
You then decide to move to square 36, ending the game.
It can be shown that you need at least 4 moves to reach the N*N-th square, so the answer is 4.
Note:

2 <= board.length = board[0].length <= 20
board[i][j] is between 1 and N*N or is equal to -1.
The board square with number 1 has no snake or ladder.
The board square with number N*N has no snake or ladder.

solution:
simple bfs
hard part is translating x position to r and c
use visited array 
make a queue
*/

func snakesAndLadders(board [][]int) int {
    boardHeight := len(board)
    boardWidth := len(board[0])
    visited := make([]bool, boardHeight * boardWidth + 1)
    q := make(Queue, 0)
    q.Push(&Item{1, 0})
    visited[1] = true
    for !q.IsEmpty() {
        x := q.Front()
        q.Pop()
        if x.val == boardHeight * boardWidth {
            return x.depth
        }
        for i := 1; i <= 6 && x.val + i <= (boardHeight * boardWidth); i++ {
            nextVal := x.val + i
            nextR, nextC :=  TranslateMove(nextVal, boardHeight, boardWidth)
            if board[nextR][nextC] > 0 {
                nextVal = board[nextR][nextC]
            }
            if !visited[nextVal] {
                q.Push(&Item{nextVal, x.depth+1})   
                visited[nextVal] = true
            }      
        } 
    }
    return -1
}

//translate x spot to row and column in grid
func TranslateMove(val int, boardHeight int, boardWidth int) (int, int) {
    //calculate row
    r := boardHeight - val / boardWidth
    if r == boardHeight || val % boardWidth != 0 {
        r--
    }
    c := val % boardWidth
    //if dir is false start from left else from right
    if c == 0 {
        if (boardHeight - r) % 2 != 0 {
            c = boardHeight-1
        }
    } else if (boardHeight - r) % 2 == 0 {
        c = boardWidth - c 
    } else {
        c = c-1
    }
    return r, c   
}

type Item struct {
    val int
    depth int
}
type Queue []*Item

func (this *Queue) Front() *Item {
    if len(*this) == 0 {
        return nil
    }
    return (*this)[0]
}

func (this *Queue) Pop() {
     if len(*this) == 0 {
        return 
    }
    (*this) = (*this)[1:]
}

func (this *Queue) Push(k *Item) {
    (*this) = append(*this, k)
}

func (this *Queue) IsEmpty() bool {
    return len(*this) == 0
}

/******dont need this anymore****/
//translate r,c to a position x
func TranslateDirection(r int, c int, boardHeight int, boardWidth int) int {
  
    val := (boardHeight-1-r) * boardWidth
    if boardHeight % 2 == 0 {
        if r % 2 == 0 {
            val = val + boardWidth-c
        } else {
            val = val+c+1
        }
    } else {
        if r % 2 == 0 {
            val = val+c+1
        } else {
            val = val + boardWidth-c
        }
    }
    return val
}
