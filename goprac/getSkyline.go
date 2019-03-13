func getSkyline(buildings [][]int) [][]int {
    return getSkylineHelper(&buildings,0,len(buildings)-1)
    
}

func getSkylineHelper(buildings *[][]int, start int, end int) [][]int {
    
    newCoordinates := make([][]int, 0) 
    
     if end - start == 0 {
            newCoordinate1 := make([]int, 0, 2)
            newCoordinate2 := make([]int, 0, 2)
            newCoordinate1 = append(newCoordinate1, (*buildings)[start][0], (*buildings)[start][2])
            newCoordinate2 = append(newCoordinate2, (*buildings)[start][1], 0)
            newCoordinates = [][]int{newCoordinate1, newCoordinate2}
            //fmt.Println(newCoordinates)
            return newCoordinates
        } 
    
    if start < end {
        mid := start + (end-start)/2
        L := getSkylineHelper(buildings, start, mid)
        R := getSkylineHelper(buildings, mid+1, end)
        newCoordinates = merge(&L, &R, buildings, start, end)
    }
    return newCoordinates     
}

func merge(L *[][]int, R *[][]int, buildings *[][]int, start, end int) [][]int {
  
    h1 := 0
    h2 := 0
    newHeight := 0
    indexL := 0
    indexR := 0
    indexN := 0
    var newCoordinates [][]int
     
    newCoordinates = make([][]int, 0, len(*R) + len(*L))
    for indexL < len(*L) && indexR < len(*R) {
        if (*L)[indexL][0] < (*R)[indexR][0] {
            h1 = (*L)[indexL][1]
            newHeight = max(h1,h2)
            appendCoordinate(&newCoordinates, &indexN, []int{(*L)[indexL][0], newHeight})
            indexL++
        } else if (*L)[indexL][0] > (*R)[indexR][0]{
            h2 = (*R)[indexR][1]
            newHeight = max(h1,h2)
            appendCoordinate(&newCoordinates, &indexN, []int{(*R)[indexR][0], newHeight})
            indexR++
        } else {
            h1 = (*L)[indexL][1]
            h2 = (*R)[indexR][1]
            newHeight = max(h1,h2)
            appendCoordinate(&newCoordinates, &indexN, []int{(*L)[indexL][0], newHeight})
            indexR++
            indexL++
        }
    
    }

    for indexL < len(*L) {
       appendCoordinate(&newCoordinates,&indexN, (*L)[indexL])
       indexL++
    }
    
    for indexR < len(*R) {
       appendCoordinate(&newCoordinates,&indexN, (*R)[indexR])
       indexR++
    }
    return newCoordinates
    
}

func appendCoordinate(newCoordinates *[][]int, indexN *int, coordinate []int) {
    //fmt.Println("appending ",coordinate, " to ",*newCoordinates)
    if (*indexN) > 0 && (*newCoordinates)[(*indexN)-1][1] == coordinate[1] {
        return
    }
    
     if (*indexN) > 0 && (*newCoordinates)[(*indexN)-1][0] == coordinate[0] {
         (*newCoordinates)[(*indexN)-1][1] = min((*newCoordinates)[(*indexN)-1][1],coordinate[1])
         return
    }
    (*newCoordinates) = append((*newCoordinates), coordinate)
    (*indexN)++
    
}

func max(a int, b int) int {
	c := b
	if a > b {
		c = a
	}
	return c
}

func min(a int, b int) int {
	c := b
	if a < b {
		c = a
	}
	return c
}
