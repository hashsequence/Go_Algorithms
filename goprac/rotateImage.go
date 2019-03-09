func rotateImage(a [][]int) [][]int {
    if len(a) != len(a[0]) {
        return a
    }    
    helperRotate(&a, 0, len(a)-1)
    return a
    
}

func helperRotate(a *[][]int, start int, end int) {
    if start >= end {
        return 
    } else {
        var temp1 int
        var temp2 int
        length := end - start
        
        for i := 0; i < length; i++ {
            temp1 = (*a)[i+start][length+start]
            (*a)[i+start][length+start] = (*a)[start][start+i]
            temp2 = (*a)[length+start][length-i+start]
            (*a)[length+start][length-i+start] = temp1
            temp1 = (*a)[length-i+start][start]
            (*a)[length-i+start][start] = temp2
            (*a)[start][start+i] = temp1
        }
        helperRotate(a,start+1, end-1)
    }
}

            
