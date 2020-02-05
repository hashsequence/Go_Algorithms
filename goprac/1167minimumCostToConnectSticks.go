/*
1167. Minimum Cost to Connect Sticks
Medium
 
You have some sticks with positive integer lengths.

You can connect any two sticks of lengths X and Y into one stick by paying a cost of X + Y.  You perform this action until there is one stick remaining.

Return the minimum cost of connecting all the given sticks into one stick in this way.

 

Example 1:

Input: sticks = [2,4,3]
Output: 14
Example 2:

Input: sticks = [1,8,3,5]
Output: 30
 

Constraints:

1 <= sticks.length <= 10^4
1 <= sticks[i] <= 10^4

solution:
use minHeap 

*/

func connectSticks(sticks []int) int {
    mh := NewMinHeap()
    for i, _ := range sticks {
        mh.InsertKey(&Item{sticks[i]})
    }
    totalCost := 0 
    for mh.size > 0 {
        first := mh.ExtractMin()
        second := mh.ExtractMin()
        if second != nil {
            totalCost += first.val + second.val
            mh.InsertKey(&Item{first.val + second.val})
        } else {
            break
        }
    }
    return totalCost
}
type Item struct {
    val int
}

type MinHeap struct {
    arr []*Item
    size int
}

func NewMinHeap() MinHeap {
    return MinHeap{make([]*Item,0), 0}
}

func Parent(i int) int {
    return (i-1)/2
}

func Left(i int) int {
    return 2*i + 1
}

func Right(i int) int {
    return 2*i + 2
}


func Swap(a, b **Item) {
    *a, *b = *b, *a
}

func (this *MinHeap) MinHeapify(root int) {
    left := Left(root)
    right := Right(root)
    smallest := root
    
    if left < this.size && this.arr[left].val < this.arr[smallest].val {
        smallest = left
    }
    
    if right < this.size && this.arr[right].val < this.arr[smallest].val {
        smallest = right
    }
    
    if smallest != root {
        Swap(&this.arr[smallest], & this.arr[root])
        this.MinHeapify(smallest)
    }
}

func (this *MinHeap) ExtractMin() *Item {
    if this.size == 0 {
        return nil
    }
    currMin := this.arr[0]
    this.size--
    this.arr[0] = this.arr[this.size]
    this.arr[this.size] = nil
    this.arr = this.arr[0:this.size]
    if this.size == 0 {
        return currMin
    }
    this.MinHeapify(0)
    return currMin
}

func (this *MinHeap) InsertKey(item *Item) {
    this.arr = append(this.arr, item)
    this.size++
    for i := this.size - 1; i >= 0 && this.arr[i].val < this.arr[Parent(i)].val; i = Parent(i) {
        Swap(&this.arr[i], &this.arr[Parent(i)])
    }
}

func (this *MinHeap) GetMin() *Item {
    if this.size == 0 {
        return nil
    }
    return this.arr[0]
}
