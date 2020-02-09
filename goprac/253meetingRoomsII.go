/*
253. Meeting Rooms II
Medium

2123

31

Add to List

Share
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.

Example 1:

Input: [[0, 30],[5, 10],[15, 20]]
Output: 2
Example 2:

Input: [[7,10],[2,4]]
Output: 1
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.


solution:
we need to find a way to know which meetings can share a room
meeting1 and meeting2 can share a room if meeting1's start time is >  meeting2 end time

if we take the earliest start time meetings and compare it with the earliest end time meetings
we can figure out which meetings can share a room at once

the idea is to imagine everything happening in real time from earliest starttime to the end,
if we know that another meeting is ending before the current earliest meeting can start, we can use the room after the previous one
ends, otherwise we need to use another room
*/

func minMeetingRooms(intervals [][]int) int {
    // if the there is only 1 or less meetings just return the number of meetings since we dont need to share
    if len(intervals) <= 1 {
        return len(intervals)
    }
    //sort the intervals from earliest to latest start times
    QSort(intervals)
    mh := NewMinHeap()
    
    //insert the first meeting on the list
    mh.InsertKey(intervals[0][1])
    
    //iterate through the list of meetings from earliest to latest starttime
    for i := 1; i < len(intervals); i++ {
        
        //if the current earliest meeting starts after the meeting that end the earliest, then we know that they can share rooms
        if intervals[i][0] >= mh.GetMin() {
            //this meeting is done so we can pop it out of the minheap
            mh.ExtractMin()
        }
        
        //insert this new meeting into the minheap
        mh.InsertKey(intervals[i][1])
    }
    
    //the minheap size is the number of rooms we need, since we are unable to pop
    return mh.size
    
    
    
}

func QSort(arr [][]int) {
    if len(arr) < 2 {
        return
    }
    left := 0
    right := len(arr) - 1
    pivot := right/2
    
    arr[right], arr[pivot] = arr[pivot], arr[right]
    
    for i, _ := range arr{
        if arr[i][0] < arr[right][0] {
            arr[i], arr[left] = arr[left], arr[i]
            left++
        }
    }
    arr[left], arr[right] = arr[right], arr[left]
    pivot = left
    QSort(arr[:pivot])
    QSort(arr[pivot+1:])    
}

func Max(a, b int) int {
    c := a 
    if b > a {
        c = b
    }
    return c
}

type MinHeap struct {
    arr []int
    size int
}

func Parent(i int) int {
    return (i-1)/2
}

func Left(i int) int {
    return 2 * i + 1
}


func Right(i int) int {
    return 2 * i + 2
}

func Swap(a, b *int) {
    *a, *b = *b, *a
}

func NewMinHeap() MinHeap {
    return MinHeap{make([]int, 0),0}
}

func (this *MinHeap) MinHeapify(root int) {
    left := Left(root)
    right := Right(root)
    smallest := root
    
    if left < this.size && this.arr[left] < this.arr[smallest] {
        smallest = left
    } 
    
    if right < this.size && this.arr[right] < this.arr[smallest] {
        smallest = right
    }
    
    if smallest != root {
        Swap(&this.arr[smallest], &this.arr[root])
        this.MinHeapify(smallest)
    }
}

func (this *MinHeap) ExtractMin() int {
    if this.size == 0 {
        return -1
    }
    
    currMin := this.arr[0]
    this.size--
    this.arr[0] = this.arr[this.size]
    this.arr[this.size] = 0
    this.arr = this.arr[:this.size]
    this.MinHeapify(0)
    return currMin
}

func (this *MinHeap) InsertKey(val int) {
    this.arr = append(this.arr, val)
    this.size++
    
    for i := this.size-1; i >= 0 && this.arr[i] < this.arr[Parent(i)]; i = Parent(i) {
        Swap(&this.arr[i], &this.arr[Parent(i)])
    }
}

func (this* MinHeap) GetMin() int {
    if this.size > 0 {
        return this.arr[0]
    }
    return -1
}
