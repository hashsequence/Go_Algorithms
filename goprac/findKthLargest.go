/*
This method is mainly an optimization of method 1. Instead of using temp[] array, use Min Heap.

1) Build a Min Heap MH of the first k elements (arr[0] to arr[k-1]) of the given array. O(k)

2) For each element, after the kth element (arr[k] to arr[n-1]), compare it with root of MH.
……a) If the element is greater than the root then make it root and call heapify for MH
……b) Else ignore it.
// The step 2 is O((n-k)*logk)

3) Finally, MH has k largest elements and root of the MH is the kth largest element.

Time Complexity: O(k + (n-k)Logk) without sorted output. If sorted output is needed then O(k + (n-k)Logk + kLogk)

*/
func findKthLargest(nums []int, k int) int {
    mh := newMinHeap(k)
    for i, num := range nums {
        if i >= 0 && i <= k-1 {
            mh.insertKey(num)
        }
        if i >= k {
            if mh.harr[0] < num {
                mh.harr[0] = num
                mh.minHeapify(0)
            }
        }
        
    }
    return mh.harr[0]
    
}

type MinHeap struct{
    harr []int
    cap int
    heapSize int
}

func newMinHeap (cap int) MinHeap {
    mh := MinHeap{
        nil,
        cap,
        0 }
    mh.harr = make([]int, cap)
    return mh
}

func parent(i int) int {
    return (i-1)/2    
}

func left(i int) int {
    return (2*i)+1
}

func right(i int) int {
    return (2*i)+2
}

func (mh *MinHeap) insertKey(k int) int {

    if mh.heapSize == mh.cap {
        return -1
    }
    
    mh.harr[mh.heapSize] = k
    pos := mh.heapSize
    mh.heapSize++
    
    for ; pos != 0 && mh.harr[pos] < mh.harr[parent(pos)]; pos = parent(pos) {
        swap(&mh.harr[pos],&mh.harr[parent(pos)])
    }
    return k;
}

func (mh *MinHeap) minHeapify(i int) {
    l := left(i)
    r := right(i)
    
    smallest := i
    
    if l < mh.heapSize && mh.harr[i] > mh.harr[l] {
        smallest = l
    }
    if r < mh.heapSize && mh.harr[smallest] > mh.harr[r] {
        smallest = r
    }
    
    if smallest != i {
        swap(&mh.harr[smallest],&mh.harr[i])
        mh.minHeapify(smallest)
    }
    
}

func swap(a *int, b *int) {
    t := *a
    *a = *b
    *b = t
}
