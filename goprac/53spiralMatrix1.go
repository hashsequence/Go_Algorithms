func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := 0
    if m > 0 {
        n = len(matrix[0])
    } else {
        return []int{}
    }
    visited := make([][]bool, m)
    for i, _ := range visited {
        visited[i] = make([]bool,n)
    }
    
    currDir := RIGHT
    currCoord := Dir{0,0}
    res := make([]int,0,n*m)

    for {
       // fmt.Println(currCoord)
        if currCoord.x >= 0 && currCoord.x < n && currCoord.y >=0 && currCoord.y < m && !visited[currCoord.y][currCoord.x] {
            res = append(res, matrix[currCoord.y][currCoord.x])
            visited[currCoord.y][currCoord.x] = true
        } else {
            break
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
            if currCoord.y+1 < m && !visited[currCoord.y+1][currCoord.x]  {
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
    return res
}

const (
    LEFT = 1
    RIGHT = 2
    UP = 3
    DOWN = 4
)

type Dir struct {
    x int
    y int
}
