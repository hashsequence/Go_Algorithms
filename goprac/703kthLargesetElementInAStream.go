/*
703. Kth Largest Element in a Stream
Easy

Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

Example:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
Note: 
You may assume that nums' length ≥ k-1 and k ≥ 1.

solution:
use minheap
remember
edge cases if k < len(nums)
nums is empty
*/

type KthLargest struct {
    mh MinHeap
    kth int
}


func Constructor(k int, nums []int) KthLargest {
    mh := newMinHeap(nums, len(nums))
    for i := 0; len(nums) > k && i < len(nums)-k; i++ {
        mh.extractMin()
    }
    return KthLargest{
        mh : mh,
        kth : k,
    }
}


func (this *KthLargest) Add(val int) int {
    if this.mh.size < this.kth {
         this.mh.insertKey(val)
    } else if m, ok := this.mh.getMin(); val > m && ok{
        this.mh.extractMin()
        this.mh.insertKey(val)
    } 
    kth, _ := this.mh.getMin()
    return kth
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */


type MinHeap struct {
    arr []int
    size int
}


func newMinHeap(nums []int, size int) MinHeap {
    mh := MinHeap{}
    mh.arr = make([]int, len(nums))
    mh.size = len(nums)
    copy(mh.arr, nums)
    mh.heapifyArr()
    return mh
    
}

func (this *MinHeap) getMin() (int, bool) {
    if this.size == 0 {
        return 0, false
    }
    return this.arr[0], true
} 

func (this *MinHeap) minHeapify(i int) {
    
    leftNode := left(i)
    rightNode := right(i)
    smallest := i
    
    if leftNode < this.size && this.arr[smallest] > this.arr[leftNode] {
        smallest = leftNode
    }
    
     if rightNode < this.size && this.arr[smallest] > this.arr[rightNode] {
        smallest = rightNode
    }
    
    if smallest != i {
        swap(&(this.arr[smallest]), &(this.arr[i]))
        this.minHeapify(smallest)
    }
}

func (this *MinHeap) insertKey(val int) {
     if this.size == len(this.arr) {
         arr := make([]int,this.size*2 + 1)
         copy(arr, this.arr)
         this.arr = arr
    }
    this.arr[this.size] = val
    this.size++
    for i := this.size-1; i != 0 && this.arr[i] < this.arr[parent(i)]; i = parent(i) {
        swap(&(this.arr[i]),&(this.arr[parent(i)]))
    }
}
func (this *MinHeap) extractMin() (int, bool) {
    if this.size <= 0 {
        return 0, false
    }
    currMin := this.arr[0]
    
    this.arr[0] = this.arr[this.size-1]
    this.size--
    this.minHeapify(0)
    return currMin, true
}

func (this *MinHeap) heapifyArr() {
    for i := (this.size-2)/2; i >=0; i-- {
        this.minHeapify(i)
    }
}

func left(i int) int {
    return 2*i + 1
}

func right(i int) int {
    return 2*i + 2
}

func parent(i int) int {
    return (i-1)/2
}

func swap(a *int, b*int) {
    (*a), (*b) = (*b), (*a)
}
