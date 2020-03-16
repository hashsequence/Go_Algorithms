/*
84. Largest Rectangle in Histogram
Hard


Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

 


Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].

 


The largest rectangle is shown in the shaded area, which has area = 10 unit.

 

Example:

Input: [2,1,5,6,2,3]
Output: 10


solution:

we need to use a stack but thats not the hard part:
algo is:
make a stack
push -1 into it denoting position before index 0
iterate through heights
loop through stack only if top is not and height[i] <= height[s.Top()]
    calculate maxArea by taking max of (height[s.Top()] * (i-1 - prevIndex))
    so if i is 4 and s.Top() is 3 and prevIndex is 2 then area = 3 * (4-1-2) = 3
    
after loop check for last retangle in the stack
where i = len(heights)

*/
func largestRectangleArea(heights []int) int {
    s := make(Stack, 0)
    s.Push(-1)
    maxArea := 0
    for i := 0; i < len(heights); i++ {
        //keep popping until we encounter -1 or current stack top is greater than current height at i 
        for !s.IsEmpty() && (s.Top() != -1 && heights[s.Top()] >= heights[i]) {
            currIndex := s.Top()
            s.Pop()
            prevIndex := s.Top()
           // fmt.Println(currIndex, prevIndex, i, heights[currIndex] * (i-1 - prevIndex)) 
            maxArea = Max(maxArea, heights[currIndex] * (i-1 - prevIndex))
        }
        s.Push(i)
    }
    //checking the final rectangle
    for s.Top() != -1 {
        currIndex := s.Top()
        s.Pop()
        prevIndex := s.Top()
        maxArea = Max(maxArea, heights[currIndex] * (len(heights)-1 - prevIndex))
    }
    return maxArea
}
type Item struct {
    val int
}
type Stack []int
func (this *Stack) Top() int {
    if len(*this) == 0 {
        return -1
    }
    return (*this)[len(*this)-1]
}

func (this *Stack) Push(k int) {
    (*this) = append(*this, k)
}

func (this *Stack) Pop() {
    if len(*this) > 0 {
        (*this) = (*this)[:len(*this)-1]
    }
}

func Max(a int, b int) int {
    if a > b {
        return a
    } 
    return b
}


func Min(a int, b int) int {
    if a < b {
        return a
    } 
    return b
}


func (this *Stack) IsEmpty() bool {
    return len(*this) == 0
}
