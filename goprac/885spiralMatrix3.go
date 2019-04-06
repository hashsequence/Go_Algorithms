func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
    res := make([][]int, R*C)
    for i, _ := range res {
        res[i] = make([]int,2)
    }
 
    i := 0   
    n := 0
    for i < R*C {
        traversePerimeter(R, C, r0, c0, n, &res, &i)
       // fmt.Println(i, n)
        n++
     }
    return res
}

func traversePerimeter(R int, C int, r0 int, c0 int, n int, res *[][]int, i *int) {
    if n == 0 {
        (*res)[*i][0] = r0
        (*res)[*i][1] = c0
        //fmt.Println((*res)[*i])
        (*i)++
        return
    } 
    //traverse right edge if possible up to down
    //fmt.Println(c0+n, C)
    if c0 + n < C {
       // fmt.Println("traversing right edge")
        c := c0 + n
        start := max(0, r0 - n + 1)
        end := min(R-1, r0 + n)
        for index := start; index <= end; index++ {
            (*res)[*i][0] = index
            (*res)[*i][1] = c
            //fmt.Println((*res)[*i])
            (*i)++
        }
    } 
    //traverse down edge if possible right to left
    if r0 + n < R {
       // fmt.Println("traversing down edge")
        r := r0 + n
        start := min(C-1, c0 + n - 1)
        end := max(0, c0 - n)
        for index := start; index >= end; index-- {
            (*res)[*i][0] = r
            (*res)[*i][1] = index
            //fmt.Println((*res)[*i])
            (*i)++
        }  
    } 
    //traverse left edge if possible down to up
    if c0 - n >= 0 {
        //fmt.Println("traversing left edge")
        c := c0 - n
        start := min(R-1, r0 + n -1)
        end := max(0, r0 - n)
        for index := start; index >= end; index-- {
            (*res)[*i][0] = index
            (*res)[*i][1] = c
            //fmt.Println((*res)[*i])
            (*i)++
        }
    } 
    //traverse up edge if possible left to right
    if r0 - n >= 0 {
        //fmt.Println("traversing up edge")
        r := r0 - n
        start := max(0, c0 - n + 1)
        end := min(C-1, c0 + n)
        for index := start; index <= end; index++ {
            (*res)[*i][0] = r
            (*res)[*i][1] = index
            //fmt.Println((*res)[*i])
            (*i)++
        }
    } 
    return 
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
