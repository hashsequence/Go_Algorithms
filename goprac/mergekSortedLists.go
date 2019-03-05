/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
const MaxUint = ^uint(0)
const MinUint = 0

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func mergeKLists(lists []*ListNode) *ListNode {
	h := newHeap(len(lists))
	var list *ListNode

	for i, currList := range lists {
			if currList != nil {
                // fmt.Println(Node{currList.Val, i}, " " , i)
				h.insertKey(Node{currList.Val, i})
				lists[i] = currList.Next
			}

		}
        
  var listIt *ListNode
    for h.heapSize != 0 {
    //    fmt.Println("current heap: ",h.harr)
		_ = listIt
		currMin := h.extractMin()
      //  fmt.Println("current min: ",currMin)

		if currMin.source != -1 {
			if list == nil {
				list = &ListNode{currMin.val, nil}
				listIt = list
			} else {
				listIt.Next = &ListNode{currMin.val, nil}
				listIt = listIt.Next
			}
            if lists[currMin.source] != nil {
              //  fmt.Println("inserting: ", lists[currMin.source].Val, " from source ", currMin.source)
			    h.insertKey(Node{(lists[currMin.source]).Val, currMin.source})
			    lists[currMin.source] = lists[currMin.source].Next
            } else {
                
            }
        } 
        
    }
		
        
    
	return list

}

type Node struct {
	val    int
	source int
}
type MinHeap struct {
	harr     []Node
	cap      int
	heapSize int
}

func newHeap(cap int) MinHeap {
	mH := MinHeap{nil, cap, 0}
	mH.harr = make([]Node, cap)
	return mH
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (2 * i) + 1
}

func right(i int) int {
	return (2 * i) + 2
}

func (this *MinHeap) insertKey(k Node) int {
	if this.cap == this.heapSize {
		return -1
	}

	pos := this.heapSize
	this.harr[this.heapSize] = k
	this.heapSize++

	for ; pos != 0 && this.harr[pos].val < this.harr[parent(pos)].val; pos = parent(pos) {
		swap(&this.harr[pos], &this.harr[parent(pos)])
	}
	return k.val
}

func (this *MinHeap) minHeapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i

	if l < this.heapSize && this.harr[l].val < this.harr[smallest].val {
		smallest = l
	}
	if r < this.heapSize && this.harr[r].val < this.harr[smallest].val {
		smallest = r
	}

	if smallest != i {
		swap(&this.harr[smallest], &this.harr[i])
		this.minHeapify(smallest)
	}
}

func (this *MinHeap) extractMin() Node {
	if this.heapSize == 0 {
		return Node{MaxInt, -1}
	}
	if this.heapSize == 1 {
		this.heapSize--
		return this.harr[0]
	}

	min := this.harr[0]
	this.harr[0] = this.harr[this.heapSize-1]
	this.heapSize--

	this.minHeapify(0)
	return min
}

func swap(a *Node, b *Node) {
	t := *a
	*a = *b
	*b = t
}
