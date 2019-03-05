/*
min max binary heap
-----------------------------------------
level order travsersal algo:
printLevelorder(tree)
1) Create an empty queue q
2) temp_node = root /*start from root
3) Loop while temp_node is not NULL
    a) print temp_node->data.
    b) Enqueue temp_node’s children (first left then right children) to q
    c) Dequeue a node from q and assign it’s value to temp_node


How is a Binary Heap represented?
A Binary Heap is a Complete Tree. A binary heap is typically represented as an array.

The root element will be at Arr[0]
Arr[(i-1)/2] returns the parent node
Arr[(2*i)+1] returns the left child nold
Arr[(2*i) + 2] returns the right child node

The traversal method use to achieve Array representation is Level Order

1) getMini(): It returns the root element of Min Heap. Time Complexity of this operation is O(1).

2) extractMin(): Removes the minimum element from MinHeap. Time Complexity of this Operation is O(Logn)
as this operation needs to maintain the heap property (by calling heapify()) after removing root.

3) decreaseKey(): Decreases value of key. The time complexity of this operation is O(Logn).
If the decreases key value of a node is greater than the parent of the node, then we don’t need to do anything. Otherwise, we need to traverse up to fix the violated heap property.

4) insert(): Inserting a new key takes O(Logn) time. We add a new key at the end of the tree.
IF new key is greater than its parent, then we don’t need to do anything. Otherwise, we need to traverse up to fix the violated heap property.

5) delete(): Deleting a key also takes O(Logn) time. We replace the key to be deleted with
 minum infinite by calling decreaseKey(). After decreaseKey(), the minus infinite value must reach
 root, so we call extractMin() to remove the key.



*/

package main

import (
	"fmt"
	"reflect"
)

const MaxUint = ^uint(0)
const MinUint = 0

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type MinHeap struct {
	harr      *[]int
	capacity  int
	heap_size int
}

func (minH *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (minH *MinHeap) left(i int) int {
	return (2 * i) + 1
}

func (minH *MinHeap) right(i int) int {
	return (2 * i) + 2
}

func (minH *MinHeap) getMin() int {
	return (*minH.harr)[0]
}

func NewMinHeap(cap int) MinHeap {
	minH := MinHeap{nil, cap, 0}
	tempArr := make([]int, cap)
	minH.harr = &tempArr
	return minH
}

func (minH *MinHeap) insertKey(k int) int {
	if minH.heap_size == minH.capacity {
		return -1
	}

	//insert key at the end
	(*minH.harr)[minH.heap_size] = k
	i := minH.heap_size
	minH.heap_size++

	//fix the min heap property if it is violated
	for ; i != 0 && (*minH.harr)[minH.parent(i)] > (*minH.harr)[i]; i = minH.parent(i) {
		swap(&(*minH.harr)[i], &(*minH.harr)[minH.parent(i)])
	}
	return k
}

func (minH *MinHeap) decreaseKey(i int, new_val int) {
	(*minH.harr)[i] = new_val
	for ; i != 0 && (*minH.harr)[minH.parent(i)] > (*minH.harr)[i]; i = minH.parent(i) {
		swap(&(*minH.harr)[i], &(*minH.harr)[minH.parent(i)])
	}
}

func (minH *MinHeap) extractMin() int {
	if minH.heap_size <= 0 {
		return MaxInt
	}
	if minH.heap_size == 1 {
		minH.heap_size--
		return (*minH.harr)[0]
	}

	root := (*minH.harr)[0]
	(*minH.harr)[0] = (*minH.harr)[minH.heap_size-1]
	minH.heap_size--
	minH.minHeapify(0)
	return root
}

func (minH *MinHeap) deleteKey(i int) {
	minH.decreaseKey(i, MinInt)
	minH.extractMin()
}

func (minH *MinHeap) minHeapify(i int) {
	l := minH.left(i)
	r := minH.right(i)

	smallest := i

	if l < minH.heap_size && (*minH.harr)[l] < (*minH.harr)[i] {
		smallest = l
	}
	if r < minH.heap_size && (*minH.harr)[r] < (*minH.harr)[smallest] {
		smallest = r
	}
	if smallest != i {
		swap(&(*minH.harr)[i], &(*minH.harr)[smallest])
		minH.minHeapify(smallest)
	}

}

func swap(a *int, b *int) bool {
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		t := *a
		*a = *b
		*b = t
		return true
	}
	return false
}

func main() {
	m := NewMinHeap(5)
	fmt.Println((*m.harr)[2])

	h := NewMinHeap(11)
	h.insertKey(3)
	h.insertKey(2)
	h.deleteKey(1)
	h.insertKey(15)
	h.insertKey(5)
	h.insertKey(4)
	h.insertKey(45)
	fmt.Println(h.extractMin())
	fmt.Println(h.getMin())
	h.decreaseKey(2, 1)
	fmt.Println(h.getMin())

}
