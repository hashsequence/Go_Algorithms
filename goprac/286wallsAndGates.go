/*
286. Walls and Gates
Medium

You are given a m x n 2D grid initialized with these three possible values.

-1 - A wall or an obstacle.
0 - A gate.
INF - Infinity means an empty room. We use the value 231 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647.
Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with INF.

Example: 

Given the 2D grid:

INF  -1  0  INF
INF INF INF  -1
INF  -1 INF  -1
  0  -1 INF INF
After running your function, the 2D grid should be:

  3  -1   0   1
  2   2   1  -1
  1  -1   2  -1
  0  -1   3   4
  
  solution;
  use bfs with a Queue
  check if the current level of the current path is less than the set value of rooms[i][j]
  and set it to the min if so
*/
func wallsAndGates(rooms [][]int)  {
    q := make(Queue,0)
    for i := 0; i < len(rooms); i++ {
        for j := 0; j < len(rooms[i]); j++ {
            if rooms[i][j] == 0 {
                q.Push(&Item{i,j, 0})
            }
        }
    }
    Dfs(rooms, q)
}

type Item struct {
    i int
    j int
    level int
}

type Queue []*Item

func (q *Queue) Push(item *Item) {
    if item == nil {
        return
    }
    *q = append(*q, item)
}

func (q *Queue) Pop() {
    if len(*q) == 0 {
        return
    }
    (*q)[0] = nil
    *q = (*q)[1:]
}

func (q *Queue) Top() *Item {
    if len(*q) == 0 {
        return nil
    }
    return (*q)[0]
}

func Dfs(rooms [][]int, q Queue) {
    
    for len(q) > 0 {
        curr := q.Top()
        q.Pop()
        i := curr.i
        j := curr.j
        level := curr.level
        rooms[i][j] = Min(rooms[i][j], level)
        nextLevel := level+1
        if i-1 >= 0 && (rooms[i-1][j] == 2147483647) {
            q.Push(&Item{i-1,j,nextLevel})
        }
        if j-1 >= 0 && (rooms[i][j-1] == 2147483647) {
            q.Push(&Item{i,j-1,nextLevel})
        }
        if i+1 < len(rooms) && (rooms[i+1][j] == 2147483647) {
            q.Push(&Item{i+1,j,nextLevel})
        }
        if j+1 < len(rooms[i]) && (rooms[i][j+1] == 2147483647) {
            q.Push(&Item{i,j+1,nextLevel})
        }
    }
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
