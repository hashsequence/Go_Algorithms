/*
1102. Path With Maximum Minimum Value
Medium

Given a matrix of integers A with R rows and C columns, find the maximum score of a path starting at [0,0] and ending at [R-1,C-1].

The score of a path is the minimum value in that path.  For example, the value of the path 8 →  4 →  5 →  9 is 4.

A path moves some number of times from one visited cell to any neighbouring unvisited cell in one of the 4 cardinal directions (north, east, west, south).

 

Example 1:



Input: [[5,4,5],[1,2,6],[7,4,6]]
Output: 4
Explanation: 
The path with the maximum score is highlighted in yellow. 
Example 2:



Input: [[2,2,1,2,2,2],[1,2,2,2,1,2]]
Output: 2
Example 3:



Input: [[3,4,6,3,4],[0,2,1,1,7],[8,8,3,2,7],[3,2,4,9,8],[4,1,2,0,0],[4,6,5,4,3]]
Output: 3
 

Note:

1 <= R, C <= 100
0 <= A[i][j] <= 10^9

solution:
the score is the minimum value in the path,
and you want to find the path with the biggest minimum 
to do this you want to pick the biggest number at each turn
use maxheap to always pick the largest node to traverse first
you must use bfs to traverse and we can use a maxHeap which is a priority queue
bfs and priority queue is one way to implement dijkstra algo

we can use dfs but is slower because
*/

func maximumMinimumPath(A [][]int) int {
    //for i,_ := range A {
    //    for j,_ := range A[i] {
    //        fmt.Print(A[i][j])
    //    }
    //    fmt.Print("\n")
    //}
    //return dfs(A, 0, 0) 
    if len(A) == 0 {
        return 0
    }
    if len(A) == 1 && len(A[0]) == 1 {
        return A[0][0]
    }
    
    maxHeap := NewMaxHeap()
    maxHeap.InsertKey(&Item{A[0][0], 0, 0})
    min := Min(A[0][0], A[len(A)-1][len(A[0])-1])
    A[0][0] = -1
    
    for maxHeap.size > 0 { 
        currNode := maxHeap.ExtractMax()
        A[currNode.i][currNode.j] = -1
        //for i,_ := range A {
        //  for j,_ := range A[i] {
        //      fmt.Print(A[i][j])
        //  }
        //  fmt.Print("\n")
        //}
        //fmt.Println(currNode.val, min)
        min = Min(currNode.val, min)
    
        dir := [4][2]int{{-1,0},{0,-1}, {1,0}, {0,1}}
        for i, _ := range dir {
            nextI := currNode.i + dir[i][0]
            nextJ := currNode.j + dir[i][1]
             
            if nextI == len(A)-1 && nextJ ==len(A[0])-1 {
            //fmt.Println("----------end----------")
                min = Min(A[nextI][nextJ], min)
                return min
            }
            
            if nextI < len(A) && nextI >= 0 && nextJ < len(A[0]) && nextJ >= 0 && A[nextI][nextJ] != -1 {
                maxHeap.InsertKey(&Item{A[nextI][nextJ], nextI, nextJ})
            }
        }
    }
    return min
}

func dfs(A [][]int, i, j int) int {
    if i < len(A) && i >= 0 && j < len(A[0]) && j >= 0 {
        if i == len(A)-1 && j ==len(A[0])-1 {
        //    fmt.Println("----------end----------")
            return A[i][j]
        }
        maxHeap := NewMaxHeap()
        if i-1 >= 0 && A[i-1][j] != -1 {
            maxHeap.InsertKey(&Item{A[i-1][j], i-1, j})
        }
        if j-1 >= 0 && A[i][j-1] != -1 {
            maxHeap.InsertKey(&Item{A[i][j-1], i, j-1})
        }
        if i+1 < len(A) && A[i+1][j] != -1 {
            maxHeap.InsertKey(&Item{A[i+1][j], i+1, j})
        } 
        if j+1 < len(A[0]) && A[i][j+1] != -1 {
            maxHeap.InsertKey(&Item{A[i][j+1], i, j+1})
        } 
        if maxHeap.size == 0 {
            return -1
        }
        score := A[i][j]
        A[i][j] = -1
        var nextNode *Item
        nextScore := -1
        maxNextScore := -1
        for maxHeap.size > 0 {
            nextNode = maxHeap.ExtractMax()
          //  fmt.Println("nextNode: ",nextNode, "\n++++++")   
            nextScore = dfs(A, nextNode.i, nextNode.j)
            if maxNextScore < nextScore {
                maxNextScore = nextScore
            }
        }
        A[i][j] = score
        if score < maxNextScore {
                return score
        }
        return maxNextScore

    }
    return -1
}


type Item struct {
    val int
    i int
    j int
}

type MaxHeap struct {
    arr []*Item
    size int
}

func NewMaxHeap() MaxHeap {
    return MaxHeap{make([]*Item, 0), 0}
}

func Parent(i int) int {
    return (i-1)/2
}

func Left(i int) int {
    return (2*i) + 1
}

func Right(i int) int {
    return (2*i) + 2
}

func Swap(a, b **Item) {
    *a, *b = *b, *a
}

func (this *MaxHeap) GetMin() *Item {
    if this.size == 0 {
        return nil
    }
    return this.arr[0]
}

func (this *MaxHeap) MaxHeapify(root int) {
    left := Left(root)
    right := Right(root)
    biggest := root
    if left < this.size && this.arr[biggest].val < this.arr[left].val {
        biggest = left
    }
    if right < this.size && this.arr[biggest].val < this.arr[right].val {
        biggest = right
    }
    if biggest != root {
        Swap(&this.arr[biggest], &this.arr[root])
        this.MaxHeapify(biggest)
    }  
}

func (this *MaxHeap) ExtractMax() *Item{
    if this.size == 0 {
        return nil
    }
    currMin := this.arr[0]
    this.arr[0] = this.arr[this.size-1]
    this.arr[this.size-1] = nil
    this.size--
    this.arr = this.arr[:this.size]
    this.MaxHeapify(0)
    return currMin
}

func (this *MaxHeap) InsertKey(item *Item) {
    this.arr = append(this.arr, item)
    this.size++
    for i := this.size-1; i > 0 &&  this.arr[i].val > this.arr[Parent(i)].val; i = Parent(i) {
        Swap(&this.arr[i], &this.arr[Parent(i)])
    }
}

func Min(a , b int) int {
    c := a 
    if b < a {
        c = b
    }
    return c
}
