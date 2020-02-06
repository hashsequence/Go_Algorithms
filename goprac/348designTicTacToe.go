/*
348. Design Tic-Tac-Toe
Medium

601

41

Add to List

Share
Design a Tic-tac-toe game that is played between two players on a n x n grid.

You may assume the following rules:

A move is guaranteed to be valid and is placed on an empty block.
Once a winning condition is reached, no more moves is allowed.
A player who succeeds in placing n of their marks in a horizontal, vertical, or diagonal row wins the game.
Example:
Given n = 3, assume that player 1 is "X" and player 2 is "O" in the board.

TicTacToe toe = new TicTacToe(3);

toe.move(0, 0, 1); -> Returns 0 (no one wins)
|X| | |
| | | |    // Player 1 makes a move at (0, 0).
| | | |

toe.move(0, 2, 2); -> Returns 0 (no one wins)
|X| |O|
| | | |    // Player 2 makes a move at (0, 2).
| | | |

toe.move(2, 2, 1); -> Returns 0 (no one wins)
|X| |O|
| | | |    // Player 1 makes a move at (2, 2).
| | |X|

toe.move(1, 1, 2); -> Returns 0 (no one wins)
|X| |O|
| |O| |    // Player 2 makes a move at (1, 1).
| | |X|

toe.move(2, 0, 1); -> Returns 0 (no one wins)
|X| |O|
| |O| |    // Player 1 makes a move at (2, 0).
|X| |X|

toe.move(1, 0, 2); -> Returns 0 (no one wins)
|X| |O|
|O|O| |    // Player 2 makes a move at (1, 0).
|X| |X|

toe.move(2, 1, 1); -> Returns 1 (player 1 wins)
|X| |O|
|O|O| |    // Player 1 makes a move at (2, 1).
|X|X|X|
Follow up:
Could you do better than O(n2) per move() operation?

solution:
to check if win
just check horizontal, vertical, diagonals

*/
/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
/*
solution with
O(N) time
O(N) space 
*/
type TicTacToeV1 struct {
    grid [][]int
    isWin int

}


/** Initialize your data structure here. */
func ConstructorV1(n int) TicTacToeV1 {
    grid := make([][]int, n)
    for i,_ := range grid {
        grid[i] = make([]int, n)
    }
    isWin := 0
    return TicTacToeV1{grid, isWin}
}


/** Player {player} makes a move at ({row}, {col}).
        @param row The row of the board.
        @param col The column of the board.
        @param player The player, can be either 1 or 2.
        @return The current winning condition, can be either:
                0: No one wins.
                1: Player 1 wins.
                2: Player 2 wins. */
func (this *TicTacToeV1) Move(row int, col int, player int) int {
    this.grid[row][col] = player
    this.isWin = this.Win(row, col, player)
    //fmt.Println(this.grid)
    return this.isWin
    
}

func (this *TicTacToeV1) Win(row, col int, player int) int {
    //check horizontal 
    horizontal, vertical, diagonal1, diagonal2 := true, true, true, true
    //check horizontal 
    for i := 0; i < len(this.grid); i++ {
        if this.grid[row][i] != player {
           horizontal = false 
           break
        } 
    }
    if horizontal {
        return player
    }
    //check vertical
    for i := 0; i < len(this.grid); i++ {
        if this.grid[i][col] != player {
           vertical = false 
            break
        }
    }
    if vertical {
        return player
    }
    //check diagonal 1
    for i, j := 0, 0; i < len(this.grid) || j < len(this.grid); i, j = i+1, j+1 {
        if this.grid[i][j] != player {
            diagonal1 = false 
            break
        }
    }
    if diagonal1 {
        return player
    }
    //check diagonal 2
    for i, j := len(this.grid)-1, 0; i >= 0 || j < len(this.grid); i, j = i-1, j+1 {
        if this.grid[i][j] != player {
            diagonal2 = false 
            break
        }
    }
    if diagonal2 {
        return player
    }
    return 0
    
}
/***********
O(1) time complexity version
O(N) space

make a grid to keep track of scores for each player
grid[player][score]
there are n rows, n cols, and 2 diagonals possible ways to win
**********/

type TicTacToe struct {
    grid [2][]int
    n int

}


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
   grid := [2][]int{}
    for i, _ := range grid {
        grid[i] = make([]int, 2*n + 2)
    }
    return TicTacToe{grid, n}
    
}

func (this *TicTacToe) Move(row int, col int, player int) int {
    //horiz
    this.grid[player-1][row]++
    if this.grid[player-1][row] == this.n {
        return player
    }
    //vertical
    this.grid[player-1][this.n + col]++
    if this.grid[player-1][this.n + col] == this.n {
        return player
    }
    //diagonal1
    if row == col {
        this.grid[player-1][this.n * 2]++
    }
    if this.grid[player-1][this.n * 2] == this.n {
        return player
    }
    //diagonal2
    if col + row + 1 == this.n {
         this.grid[player-1][this.n * 2 + 1]++
    }
    if this.grid[player-1][this.n * 2 + 1] == this.n {
        return player
    }
    return 0
}
