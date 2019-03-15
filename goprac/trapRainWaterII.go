func trapRainWater(heightMap [][]int) int {
    if len(heightMap) < 3 {
        return 0
    }
    if heightMap != nil {
        if len(heightMap[0]) < 3 {
            return 0
        }
    }
    water := 0
    
    return water
}

const MaxUint = ^uint(0)
const MinUint = 0

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type Cell struct {
    x int
    y int
    height int
}

func (this *Cell) compare(c *Cell) int {
    if this.height < c.height {
        return -1
    } else if this.height > c.height {
        return 1
    } else {
        return 0
    }
}
type MinHeap struct {
    arr *[]Cell
    cap  int
	size int
}

func NewMinHeap(cap int) MinHeap {
	minH := MinHeap{nil, cap, 0}
	tempArr := make([]Cell, cap)
	minH.arr = &tempArr
	return minH
}

func (this *MinHeap) parent(i int) int {
    return (i-1)/2
}

func (this *MinHeap) left(i int) int {
    return 2*i+1
}

func (this *MinHeap) right(i int) int {
    return 2*i+2
}

func (this *MinHeap) minHeapify(i int) {
    if this.cap == this.size {
        return
    } 
    
    l := this.left(i)
    r := this.right(i)
    smallest := i
    
    if l < this.size && (*this.arr)[l].compare(&(*this.arr)[smallest]) == -1 {
        smallest = l
    }
    
    if r < this.size && (*this.arr)[r].compare(&(*this.arr)[smallest]) == -1 {
        smallest = r
    }
    
    if smallest != i {
        swap(&(*this.arr)[i], &(*this.arr)[smallest])
        this.minHeapify(smallest)
    }
}

func (this *MinHeap) insertKey(k Cell) {
	if this.size == this.cap {
		return 
	}

	//insert key at the end
	(*this.arr)[this.size] = k
	i := this.size
	this.size++

	//fix the min heap property if it is violated
    for ; i != 0 && (*this.arr)[this.parent(i)].compare(&(*this.arr)[i]) == 1; i = this.parent(i) {
		swap(&(*this.arr)[i], &(*this.arr)[this.parent(i)])
	}
}


func swap(a *Cell, b *Cell) bool {
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		t := *a
		*a = *b
		*b = t
		return true
	}
	return false
}

func (this *MinHeap) extractMin() Cell {
	if this.size <= 0 {
        return Cell{-1,-1,-1}
	}
	if this.size == 1 {
		this.size--
		return (*this.arr)[0]
	}

	root := (*this.arr)[0]
	(*this.arr)[0] = (*this.arr)[this.size-1]
	this.size--
	this.minHeapify(0)
	return root
}
