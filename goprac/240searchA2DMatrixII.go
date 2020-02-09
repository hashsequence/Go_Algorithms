/*
240. Search a 2D Matrix II
Medium

2351

65

Add to List

Share
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
Example:

Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.

solution:

this is the O(n + m) solution 
basically I start current at the bottom left corner
if the target is less than the current then I go up a row
if target is greater than current I can go right

we can start at the upper left because the target can be in the same row or lower row

You have to see that:
remember for each matrix[i][j] anything right and and below is bigger
<cur>*******
*    *
*     *
*      *
                                     

*/
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    i := len(matrix)-1
    j := 0
    
    for i >= 0 && j < len(matrix[0]) {
        if target == matrix[i][j] {
            return true
        }
        if matrix[i][j] > target {
            i--
        } else {
            j++
        }
    }
    return false
}


// this is the O(nlogm) solution
//I can switch wether to iterate over columns and binary search instead if m < n to get O(mlogn)
func searchMatrixv2(matrix [][]int, target int) bool {
    
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    for i,_ := range matrix {
        row := BinarySearchCol(matrix, i, target)
        //fmt.Println("row", row)
        if row < 0 || row >= len(matrix) {
            continue
        }
        col:= BinarySearchRow(matrix, row, target)
        //fmt.Println("col" ,col)
        if col < 0 || col >= len(matrix[i]) {
            continue
        }
        //fmt.Println(matrix[row][col])
        if matrix[row][col] == target {
        return true
        }
    }
    return false
}


func BinarySearchRow(matrix[][]int, r, target int) int {
    if len(matrix[r]) == 1 {
        return 0
    }
    
    if len(matrix[r]) == 2 {
        if target < matrix[r][1] {
            return 0
        } 
        return 1
    }
    row := matrix[r]
    left := 0
    right := len(row)-1
    mid := right/2
    
    if target < row[left] {
        return -1
    }
    if target > row[right] {
        return right
    }
    
    for left < right {
        mid = left + (right-left)/2
        if target < row[mid] {
            right = mid-1
        } else if target > row[mid] {
            left = mid + 1
        } else {
            return mid
        }
    }
    //fmt.Println("Right", right, " left ", left)
     if row[right] > target {
        return mid
    }
    return right
    
    
}

func BinarySearchCol(matrix[][]int, c ,target int) int {
    
    if len(matrix) == 1 {
        return 0
    }
    
    if len(matrix) == 2 {
        if target < matrix[1][c] {
            return 0
        } 
        return 1
    }
    
    left := 0
    right := len(matrix)-1
    mid := 0
    
    if target < matrix[left][c] {
        return -1
    }
    if target > matrix[right][c] {
        return right
    }
    
    for left < right {
        mid = left + (right-left)/2
        if target < matrix[mid][c] {
            right = mid-1
        } else if target > matrix[mid][c] {
            left = mid + 1
        } else {
            return mid
        }
    }
    //fmt.Println("Right", right, "left", left, "mid", mid)
    if matrix[right][c] > target {
        return mid
    }
    return right
    
    
}




