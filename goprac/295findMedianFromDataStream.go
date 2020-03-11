/*
295. Find Median from Data Stream
Hard

Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

For example,
[2,3,4], the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:

void addNum(int num) - Add a integer number from the data stream to the data structure.
double findMedian() - Return the median of all elements so far.
 

Example:

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3) 
findMedian() -> 2
 

Follow up:

If all integer numbers from the stream are between 0 and 100, how would you optimize it?
If 99% of all integer numbers from the stream are between 0 and 100, how would you optimize it?

solution:
use a minHeap and a maxHeap

make a balance function to balance the heap if size difference reaches 2

if num you are adding is less than the top element in the minHeap put it into the maxHeap

time complexity: O(logn) since each operation involves popping and pushing the number O(1) times
space complexity: O(n) since we are using two heaps
*/
type MedianFinder struct {
    minHeap Heap
    maxHeap Heap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{NewHeap(Less), NewHeap(Greater)}
}


func (this *MedianFinder) AddNum(num int)  {
    if this.minHeap.size == 0 && this.maxHeap.size == 0{
        this.minHeap.Push(&Item{num})
        return
    }
    
    if num < this.minHeap.Top().val {
        this.maxHeap.Push(&Item{num})
    } else {
        this.minHeap.Push(&Item{num})
    }
    Balance(&(this.minHeap), &(this.maxHeap))
}


func (this *MedianFinder) FindMedian() float64 {
    if this.minHeap.size - this.maxHeap.size == 1{
        return float64(this.minHeap.Top().val)
    }
    if this.maxHeap.size - this.minHeap.size == 1{
        return float64(this.maxHeap.Top().val)
    } 
    if this.minHeap.size != 0 && this.maxHeap.size != 0 {
      return float64(this.minHeap.Top().val + this.maxHeap.Top().val)/2.0  
    }
    return 0
    
}

func Balance(minHeap, maxHeap *Heap) {
    if minHeap.size-maxHeap.size == 2 {
        ele := minHeap.Top()
        minHeap.Pop()
        maxHeap.Push(ele)
    } else if maxHeap.size - minHeap.size == 2 {
        ele := maxHeap.Top()
        maxHeap.Pop()
        minHeap.Push(ele)
    }
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type Item struct {
    val int    
}

func Less(o1, o2 *Item) bool {
    return o1.val < o2.val
}

func Greater(o1, o2 *Item) bool {
    return o1.val > o2.val
}

type Heap struct {
    arr []*Item
    size int
    compare func(o1, o2 *Item) bool
}

func NewHeap( c func(o1, o2 *Item) bool) Heap {
    return Heap{make([]*Item,0),0, c}
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

func (h *Heap) Heapify(root int) {
    curr := root
    l := Left(root)
    r := Right(root)
    if l < h.size &&  h.compare(h.arr[l], h.arr[curr]) {
        curr = l
    } 
    if r < h.size &&  h.compare(h.arr[r], h.arr[curr]) {
        curr = r
    } 
    if curr != root {
        Swap(&(h.arr[root]), &(h.arr[curr]))
        h.Heapify(curr)
    }
}

func (h *Heap) Top() *Item {
    if h.size > 0 {
        return h.arr[0]
    }
    return nil
}

func (h* Heap) Push(k *Item) {
    h.arr = append(h.arr, k)
    h.size++
    for i := h.size-1; i > 0 && h.compare(h.arr[i], h.arr[Parent(i)]); i = Parent(i) {
        Swap(&(h.arr[i]), &(h.arr[Parent(i)]))
    }
}

func (h* Heap) Pop() {
    if h.size == 0 {
        return
    }
    h.arr[0] = h.arr[h.size-1]
    h.arr = h.arr[:h.size-1]
    h.size--
    h.Heapify(0)
}

