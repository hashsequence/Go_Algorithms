/*
957. Prison Cells After N Days
Medium

316

538

Add to List

Share
There are 8 prison cells in a row, and each cell is either occupied or vacant.

Each day, whether the cell is occupied or vacant changes according to the following rules:

If a cell has two adjacent neighbors that are both occupied or both vacant, then the cell becomes occupied.
Otherwise, it becomes vacant.
(Note that because the prison is a row, the first and the last cells in the row can't have two adjacent neighbors.)

We describe the current state of the prison in the following way: cells[i] == 1 if the i-th cell is occupied, else cells[i] == 0.

Given the initial state of the prison, return the state of the prison after N days (and N such changes described above.)

 

Example 1:

Input: cells = [0,1,0,1,1,0,0,1], N = 7
Output: [0,0,1,1,0,0,0,0]
Explanation: 
The following table summarizes the state of the prison on each day:
Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
Day 7: [0, 0, 1, 1, 0, 0, 0, 0]

Example 2:

Input: cells = [1,0,0,1,0,0,1,0], N = 1000000000
Output: [0,0,1,1,1,1,1,0]
 

Note:

cells.length == 8
cells[i] is in {0, 1}
1 <= N <= 10^9

solution:

must implement a cycle detection, and stop iteration when we see the same same sequence again
must notice that day 0 will never repeat, so we dont need to cache it
also to get the Nth day when we see a cycle we can do (N-k) % k where k is the number of unique days,
however when (N-k) % k == 0 it is the same as day k so set the sequence for day 0 in cache equal to day k

*/
func prisonAfterNDays(cells []int, N int) []int {
    if N == 0 {
        return cells
    }
    cycleDetected := false
    set := map[[8]rune]int{}
    dayToCells := map[int][]int{}
    k := 0
    for  ;k <= N; k++ {
        //fmt.Println(k, cells)
        if _, ok := set[hash(cells)]; !ok && k != 0{
            set[hash(cells)] = k-1
            dayToCells[k] = cells
        } else {
            if k != 0 {
                cycleDetected = true
                k--
                break
            }
        }
        nextCells := make([]int,8)
        copy(nextCells, cells)
        for  i := 0; i < 8; i++ {
            if i > 0 && i < 7 && cells[i-1] == cells[i+1] {
                nextCells[i] = 1
            } else {
                nextCells[i] = 0
            }
        }
        cells = nextCells
    }
    //fmt.Println(N, k, dayToCells)
    if !cycleDetected {
        return dayToCells[N]
    }
    dayToCells[0] = dayToCells[k]
    return dayToCells[(N-k) % k]
}



func hash(cells []int) [8]rune {
    str := [8]rune{}
    for i := 0; i < 8; i++ {
        if cells[i] == 1 {
            str[i] = '1'
        } else {
            str[i] = '0'
        }
    }
    return str
}

//brute force
func prisonAfterNDaysBruteForce(cells []int, N int) []int {
    
    for k := 0; k < N; k++ {
        nextCells := make([]int,8)
        copy(nextCells, cells)
        for i := 0; i < 8; i++ {
            if i > 0 && i < 7 && cells[i-1] == cells[i+1] {
                nextCells[i] = 1
            } else {
                nextCells[i] = 0
            }
        }
        cells = nextCells
    }
    return cells
}
