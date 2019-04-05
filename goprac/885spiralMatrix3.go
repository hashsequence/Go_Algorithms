func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
    r := r0
    c := c0
    currDir := RIGHT
  
    res := make([][]int, R*C)
    for i, _ := range res {
        res[i] = make([]int,2)
    }
 
    i := 0   
    for R*C {
        res[i][0] = r
        res[i][1] = c
      
     }
}

const (
    LEFT = 1
    RIGHT = 2
    UP = 3
    DOWN = 4
)

traversePerimeter(R int, C int, r0 int, c0 int, n int, res *[][]int, i *int) {
    if n == 0 {
        res[*i][0] = r0
        res[*i][1] = c0
        return
    } 
    //traverse right edge if possible

    //traverse down edge if possible
    
    //travers left edge if possible
    
    //traverse up edge if possible
}

func min(a int, b int) int {
    c := a
    if a > b {
        c = b
    }
    return c
}

func max(a int, b int) int {
    c := a
    if a < b {
        c = b
    }
    return c
}
