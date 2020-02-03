/*
994. Rotting Oranges
Easy


In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.

solution:
iterate over grid and keep track of fresh and rotten oranges
use a queue to store the initial rotten oranges in the grid
denote 2 as starting level
then iterate over the Queue
and you evaluate all the rotten oranges on the same level
increment the level(in this case the variable is called minutes)
if there are fresh oranges left than return -1
otherwise return minutes -3, -3 comes from the fact we started at 2 and we increment after
the final level is evaluated




*/

func orangesRotting(grid [][]int) int {
    q := make(Queue,0)
    numRotten := 0
    numFresh := 0
    for i,_ := range grid {
        for j,_ := range grid[i] {
            if grid[i][j] == 2 {
                q.Push(&Node{2, i, j})
                numRotten++
            }
            if grid[i][j] == 1 {
                numFresh++
            }            
        }
    }
    if numRotten == 0 && numFresh > 0{
        return -1
    }
    if numFresh == 0 {
        return 0
    }
   
    //fmt.Println(grid)
    minutes := 2
    for len(q) > 0 {
        curr := q.Front()
        //fmt.Println(minutes)
        for curr != nil && curr.level == minutes {
            q.Pop()
            i, j := curr.i, curr.j
            if i - 1 >= 0 &&  grid[i-1][j] == 1 {
              grid[i-1][j] = minutes+1
              q.Push(&Node{minutes+1, i-1, j})
                numFresh--
            } 
            if j - 1 >= 0 &&  grid[i][j-1] == 1 {
                grid[i][j-1] = minutes+1
                q.Push(&Node{minutes+1, i, j-1})
                numFresh--
            } 
            if i + 1 < len(grid) && grid[i+1][j] == 1 {
                grid[i+1][j] = minutes+1
                q.Push(&Node{minutes+1, i+1, j})
                numFresh--
            } 
            if j + 1 < len(grid[i]) && grid[i][j+1] == 1 {
                grid[i][j+1] = minutes+1
                q.Push(&Node{minutes+1, i, j+1})
                numFresh--
            }
            if curr != nil {
                curr = q.Front()
            }
        }
        minutes++
    }
    //fmt.Println(grid)
    //fmt.Println(numFresh)
    if numFresh > 0 {
        return -1
    }
    return minutes-3
}


type Node struct {
    level int
    i int
    j int
}

type Queue []*Node

func (q *Queue) Front() *Node {
    if len(*q) > 0 {
        return (*q)[0]
    }
    return nil
}

func (q *Queue) Push(n *Node) {
    (*q) = append(*q, n)
}

func (q *Queue) Pop() {
    if len(*q) < 1 {
        *q = make([]*Node,0)
    }
    (*q)[0] = nil
    (*q) = (*q)[1:]
}

