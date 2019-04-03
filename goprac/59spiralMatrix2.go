/*
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/
func generateMatrix(n int) [][]int {
    matrix := make([][]int,n)
    for i, _ := range matrix {
        matrix[i] = make([]int,n)
    }
    visited := make([][]bool, n)
    for i, _ := range visited {
        visited[i] = make([]bool,n)
    }
    i := 1
    currDir := RIGHT
    currCoord := Dir{0,0}
    for i <= n*n {
          if currCoord.x >= 0 && currCoord.x < n && currCoord.y >=0 && currCoord.y < n && !visited[currCoord.y][currCoord.x] {
            matrix[currCoord.y][currCoord.x] = i
            i++
            visited[currCoord.y][currCoord.x] = true
        } 
        
        //fmt.Println(currDir, currCoord, visited)
        if currDir == RIGHT {
            if currCoord.x+1 < n && !visited[currCoord.y][currCoord.x+1] {
                currCoord.x++
            } else {
                currDir = DOWN
                currCoord.y++
            }
        } else if currDir == DOWN {
            if currCoord.y+1 < n && !visited[currCoord.y+1][currCoord.x]  {
                currCoord.y++
            } else {
                 currDir = LEFT
                 currCoord.x--
            }
        } else if currDir == LEFT {
            if currCoord.x-1 >= 0 && !visited[currCoord.y][currCoord.x-1] {
                currCoord.x--
            } else {
                currDir = UP
                currCoord.y--
            }
        } else if currDir == UP {
            if currCoord.y-1 >= 0 && !visited[currCoord.y-1][currCoord.x] {
                currCoord.y--
            } else {
                currDir = RIGHT
                currCoord.x++
            }
        }
    }
    return matrix
}

type Dir struct {
    x int
    y int 
}

const (
    LEFT = 1
    RIGHT = 2
    UP = 3
    DOWN = 4
)
