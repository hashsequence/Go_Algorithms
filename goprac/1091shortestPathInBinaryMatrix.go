/*
1091. Shortest Path in Binary Matrix
Medium


In an N by N square grid, each cell is either empty (0) or blocked (1).

A clear path from top-left to bottom-right has length k if and only if it is composed of cells C_1, C_2, ..., C_k such that:

Adjacent cells C_i and C_{i+1} are connected 8-directionally (ie., they are different and share an edge or corner)
C_1 is at location (0, 0) (ie. has value grid[0][0])
C_k is at location (N-1, N-1) (ie. has value grid[N-1][N-1])
If C_i is located at (r, c), then grid[r][c] is empty (ie. grid[r][c] == 0).
Return the length of the shortest such clear path from top-left to bottom-right.  If such a path does not exist, return -1.

 

Example 1:

Input: [[0,1],[1,0]]


Output: 2

Example 2:

Input: [[0,0,0],[1,1,0],[1,1,0]]


Output: 4

 

Note:

1 <= grid.length == grid[0].length <= 100
grid[r][c] is 0 or 1

solution:

I need to bfs the tree
but I need to keep track of shortest path at each coordinate

I can use a Queue for BFS

I only want add a direction to the queue if I have not traversed it to yet, and if I have, the new path
has to be a shorter path.
*/
func shortestPathBinaryMatrix(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return -1
    }
    if grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
        return -1
    }
    dirArr := [][]int{{0,-1}, {1,-1}, {1,0}, {1,1}, {0,1}, {-1,1}, {-1,0}, {-1,-1}}
    dp := make([][]int, len(grid))
    for i,_ := range dp {
        dp[i] = make([]int, len(grid[0]))
    }
    q := make(Queue,0)
    q.Push(&Dir{1,0,0})
    dp[0][0] = 1
    for len(q) > 0 {
        curr := q.Top()
        //fmt.Println(curr.depth, curr.i, curr.j, dp)
        q.Pop()
        for k,_ := range dirArr {
            i := curr.i + dirArr[k][0]
            j := curr.j + dirArr[k][1]
            depth := curr.depth+1
            if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 0 {
                 if dp[i][j] == 0 {
                    dp[i][j] = depth
                    q.Push(&Dir{depth, i,j})
                } else if depth < dp[i][j] {
                    q.Push(&Dir{depth, i,j})
                }
            }
        }
    }
    //fmt.Println(dp)
    if dp[len(grid)-1][len(grid[0])-1] == 0 {
        return -1
    }
    return dp[len(grid)-1][len(grid[0])-1]
}

type Dir struct {
    depth int
    i int
    j int
}

type Queue []*Dir 

func (q * Queue) Top() *Dir {
    if len(*q) == 0 {
       return nil
    }
    return (*q)[0]
}

func (q* Queue) Push(item *Dir) {
    (*q) = append(*q,item)
}


func (q* Queue) Pop() {
    if len(*q) > 0 {
        (*q)[0] = nil
        (*q) = (*q)[1:]
    }
 }
