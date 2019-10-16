func longestCommonSubsequence(text1 string, text2 string) int {
    
    if len(text1) < len(text2) {
        str := text1
        text1 = text2
        text2 = str
    }
    
    
    
    lp := make([][]int,len(text1))
    for i,_ := range lp {
        lp[i] = make([]int,len(text2))
    }
    for i,_ := range text1 {
        for j,_ := range text2 {
            diagonal, left, up := 0, 0, 0
            if j-1 >= 0 {
                left = lp[i][j-1]
            }
            if i-1 >= 0 {
                up = lp[i-1][j]
            }
            if i-1 >= 0 && j-1 >= 0 {
                diagonal = lp[i-1][j-1]
            }      
            lp[i][j] = func() int {
                if text1[i] == text2[j] {
                    return diagonal + 1
                } else {
                    return max(up,left)
                }
            }()
            
            
        }
    }
    //fmt.Println(lp)
    return lp[len(text1)-1][len(text2)-1]
}



func max(arr ...int) int {
    m := arr[0]
    for _,val := range arr {
        if val > m {
            m = val
        }
    }
    return m
}
