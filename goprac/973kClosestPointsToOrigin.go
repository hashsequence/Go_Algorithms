/*
973. K Closest Points to Origin
Medium

We have a list of points on the plane.  Find the K closest points to the origin (0, 0).

(Here, the distance between two points on a plane is the Euclidean distance.)

You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)

 

Example 1:

Input: points = [[1,3],[-2,2]], K = 1
Output: [[-2,2]]
Explanation: 
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
Example 2:

Input: points = [[3,3],[5,-1],[-2,4]], K = 2
Output: [[3,3],[-2,4]]
(The answer [[-2,4],[3,3]] would also be accepted.)
 

Note:

1 <= K <= points.length <= 10000
-10000 < points[i][0] < 10000
-10000 < points[i][1] < 10000

solution:
write a SumOfSquares function to calculate distance relative to the origin
use a heap to keep track of k closest

//remember how to implement minHeap

*/

func kClosest(points [][]int, K int) [][]int {
    mh := NewMinHeap(K)
    res := make([][]int, K)
    for i,_ := range points {
        item := NewItem(points[i][0],points[i][1])
        mh.InsertKey(&item)
    }
    //for i, _ := range mh.arr {
    //    fmt.Println(*mh.arr[i])
    //}
    for i := 0; i < K; i++ {
        item := mh.ExtractMin()
        res[i] = []int{item.x, item.y}
    }
    return res
}

func SumOfSquares(a, b int) int {
    return a*a + b*b
}

type Item struct {
    sum int
    x int
    y int
}

func NewItem(x, y int) Item {
    return Item{SumOfSquares(x,y),x,y}
}

type MinHeap struct {
    arr []*Item
    size int
    cap int
}

func NewMinHeap(cap int) MinHeap {
    return MinHeap{make([]*Item,0,cap), 0, cap}
}

func (this *Item) Compare(other *Item) int {
    if this.sum < other.sum {
        return -1
    } else if this.sum > other.sum {
        return 1
    }
    return 0
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

func (this *MinHeap) Heapify(root int) {
    smallest := root
    left := Left(root)
    right := Right(root)
    if left < this.size && this.arr[left].Compare(this.arr[smallest]) == -1 {
        smallest = left
    }
    if right < this.size && this.arr[right].Compare(this.arr[smallest]) == -1 {
        smallest = right
    }
    if root != smallest {
        Swap(&this.arr[smallest],&this.arr[root])
        this.Heapify(smallest)
    }
}

func (this *MinHeap) GetMin() *Item {
    if this.size > 0 {
         return this.arr[0]
    } 
    return nil
   
}

func (this *MinHeap) ExtractMin() *Item {
    if this.size <= 0 {
        return nil
    }
    val := this.arr[0]
    this.size--
    if this.size == 0 {
        this.arr[0] = nil
        return val
    }
    Swap(&this.arr[0],&this.arr[this.size])
    this.Heapify(0)
    return val
    
}

func (this *MinHeap) InsertKey(item *Item) {
    if this.size == this.cap || this.size == 0{
        heapArr := make([]*Item, this.size, 2 * this.cap + 1)
        copy(heapArr, this.arr)
        this.arr = heapArr
    }
    this.arr = append(this.arr,item)
    this.size++
    for i := this.size-1;i != 0 && this.arr[i].Compare(this.arr[Parent(i)]) == -1; i = Parent(i) {
        Swap(&this.arr[i], &this.arr[Parent(i)])
    }
}

func Swap(a, b **Item) {
    *a, *b = *b, *a
}
