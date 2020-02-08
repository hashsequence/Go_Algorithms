/*
200. Number of Islands
Medium

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input:
11110
11010
11000
00000

Output: 1
Example 2:

Input:
11000
11000
00100
00011

Output: 3

solution:
iterate over grid:
everytine you encounter a 1 increment numIslands and bfs using stack to flip the 1's to 0's so we wont check again

*/

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    numIslands := 0
    for i, _ := range grid {
        for j,_ := range grid[i] {
            if grid[i][j] == '1' {
                numIslands++
                Dfs(grid, i, j)
               // fmt.Println(grid)
            }
        }
    }
    return numIslands
}

func Dfs(grid [][]byte, i, j int) {
    s := make(Stack,0)
    s.Push(&Dir{i, j})
    for len(s) > 0 {
        curr := s.Front()
       // fmt.Println(curr.i,curr.j)
        s.Pop()
        grid[curr.i][curr.j] = '0'
        i = curr.i
        j = curr.j
        if i-1 >=0 && grid[i-1][j] == '1' {
            s.Push(&Dir{i-1,j})
        }
        if j-1 >=0  && grid[i][j-1] == '1' {
            s.Push(&Dir{i,j-1})
        }
        if i+1 < len(grid)  && grid[i+1][j] == '1' {
            s.Push(&Dir{i+1,j})
        }
        if j+1 < len(grid[0])  && grid[i][j+1] == '1' {
            s.Push(&Dir{i,j+1})
        }
    }
}

type Dir struct {
    i int
    j int
}

type Stack []*Dir

func (s *Stack) Front() *Dir {
    if len(*s) > 0 {
        return (*s)[len(*s)-1]
    }
    return nil
}

func (s *Stack) Pop() {
    if len((*s)) > 0 {
        (*s)[len(*s)-1] = nil
        (*s) = (*s)[:len((*s))-1]
    }
}

func (s *Stack) Push(dir *Dir) {
    *s = append(*s, dir)
}

