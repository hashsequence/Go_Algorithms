/*
155. Min Stack
Easy

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
 

Example:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.

solution:
O(1) for all operations
if you keep a stack of 
Item {
val int
min int
}

where min is the min at the time the item was pushed into the stack
*/

type MinStack struct {
    stack []*Item
    size int
}

type Item struct {
    val int
    min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{make([]*Item, 0), 0}
}


func (this *MinStack) Push(x int)  {
    if this.size == 0 {
        this.stack = append(this.stack,&Item{x,x})
        this.size++
        return
    }
    min := this.stack[this.size-1].min
    if min > x {
        min = x
    }
    this.stack = append(this.stack,&Item{x,min})
    this.size++
}


func (this *MinStack) Pop()  {
    if this.size == 0 {
        return 
    }
    this.size--
    this.stack[this.size] = nil
    this.stack = this.stack[:this.size]
}


func (this *MinStack) Top() int {
    if this.size == 0 {
        return 0
    }
    return this.stack[this.size-1].val
}


func (this *MinStack) GetMin() int {
    if this.size == 0 {
        return 0
    }
    return this.stack[this.size-1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
/*

//tried to write a solution with a minheap

const (
    MaxInt  = 1<<(31) - 1 // 1<<31 - 1 or 1<<63 - 1
    MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
    MaxUint = 1<<31 - 1     // 1<<32 - 1 or 1<<64 - 1
)

type MinStack struct {
    heap []*int
    stack []*int
    size int
}


func Constructor() MinStack {
    return MinStack{make([]*int,0), make([]*int,0), 0}
}

//O(logn)
func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, &x)
    this.heap = append(this.heap, &x)
    this.size++
    
    i := this.size-1
    for ;i > 0 && *(this.heap[i]) < *(this.heap[Parent(i)]); i = Parent(i) {
        Swap(&this.heap[i],&this.heap[Parent(i)])
    }
}

//O(N)
func (this *MinStack) Pop()  {
    if this.size == 0 {
        return 
    }
    this.size--
    deletedElement := this.stack[this.size] 
    this.stack = this.stack[:this.size]
    deleteIndex := 0
    for i := 0; i <= this.size; i ++ {
        if deletedElement == this.heap[i] {
            deleteIndex = i
        }
    }
    this.heap[deleteIndex] = nil
    this.heap = append(this.heap[:deleteIndex], this.heap[deleteIndex+1:]...)
    for i := (this.size-1)/2; i >= 0; i-- {
        this.MinHeapify(i)
    }
}

//O(1)
func (this *MinStack) Top() int {
    if this.size == 0 {
        return int(MinInt)
    }
    return *(this.stack[this.size-1])
}

//O(1)
func (this *MinStack) GetMin() int {
    if this.size == 0 {
        return int(MinInt)
    }
    return *(this.heap[0])
}

func Parent(i int) int {
    return (i-1)/2
}

func Left(i int) int {
    return 2*i+1
}

func Right(i int) int {
    return 2*i+2
}

func Swap(a, b **int) {
    *a, *b = *b, *a
}

func (this *MinStack) MinHeapify(root int)  int{
    left := Left(root)
    right := Right(root)
    smallest := root
    
    if left < this.size && *(this.heap[left]) < *(this.heap[smallest]) {
        smallest = left
    }
    if right < this.size && *(this.heap[right]) < *(this.heap[smallest]) {
        smallest = right
    }
    
    if smallest != root {
        Swap(&this.heap[smallest], &this.heap[root])
        return this.MinHeapify(smallest)
    }
    return smallest
}

*/
