/*
694. Number of Distinct Islands
Medium

578

39

Add to List

Share
Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Count the number of distinct islands. An island is considered to be the same as another if and only if one island can be translated (and not rotated or reflected) to equal the other.

Example 1:
11000
11000
00011
00011
Given the above grid map, return 1.
Example 2:
11011
10000
00001
11011
Given the above grid map, return 3.

Notice that:
11
1
and
 1
11
are considered different island shapes, because we do not consider reflection / rotation.
Note: The length of each dimension in the given grid does not exceed 50.
solution:
first think of a way to hash the islands:
basically I converted the coordinates relative to (0,0) as the value of a distinct islands

to iterate through the islands use dfs grid searching algo
implement Itoa to convert int to strings
*/



func numDistinctIslands(grid [][]int) int {
    fmt.Println(Itoa(0-1))
    if len(grid) <= 0 {
        return 0
    }
    set := map[string]bool{}
    for i,_ := range grid {
        for j,_ := range grid[i] {
            if grid[i][j] == 1 {
                str := new(string)
                HashIsland(i, j, i, j, grid, str)
                set[*str] = true
              
            }
              
        }
    }
    return len(set)
}

func HashIsland(relativeX, relativeY, i, j int, grid [][]int, hashStr *string)  {
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
        *hashStr = *hashStr + "(" + Itoa(i - relativeX) + "," + Itoa(j - relativeY) + ")" 
         grid[i][j] = -1
        if i-1 >= 0 && grid[i-1][j] == 1 {
            HashIsland(relativeX, relativeY, i-1, j, grid, hashStr)
        }
        if i+1 < len(grid) && grid[i+1][j] == 1 {
            HashIsland(relativeX, relativeY, i+1, j, grid, hashStr)
        } 
        if j-1 >= 0 && grid[i][j-1] == 1 {
            HashIsland(relativeX, relativeY, i, j-1, grid, hashStr)
        }
        if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
            HashIsland(relativeX, relativeY, i, j+1, grid, hashStr)
        } 
    }   
}

func Itoa(num int) string {
    if num == 0 {
        return "0"
    }
     if num == 1 {
        return "1"
    } 
     if num == -1 {
        return "-1"
    } 
    
    isNeg := false
    if num < 0 {
	    isNeg = true
        num *= -1
    }

    str := make([]byte, 0)
    for i := 1; i < num; i *= 10 {
        str = append(str, byte(((num % (i*10)) - (num % i))/i) + '0')
    }
    if isNeg {
        str = append(str, '-')
    }
    Reverse(str)
    return string(str)
}

func Reverse(s []byte) {
    l := 0
    r := len(s)-1
    for l < r {
        Swap(&s[l], &s[r])
	l++
	r--
    }
}

func Swap(a, b *byte) {
    *a, *b = *b, *a
}
