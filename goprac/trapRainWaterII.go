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
    //we want to use a minHeap 
    visited := make([][]bool,len(heightMap))
    h := newMinHeap(len(heightMap) * len(heightMap[0]))
    
    for i, _ := range visited {
        visited[i] = make([]bool, len(heightMap[i]))
    }
    
    for i := 0; i < len(heightMap[0]); i++ {
        h.insertKey(Cell{0, i, heightMap[0][i]})
        h.insertKey(Cell{len(heightMap)-1, i, heightMap[len(heightMap)-1][i]})
        visited[0][i] = true
        visited[len(heightMap)-1][i] = true
    }
    
    for i := 0; i < < len(heightMap); i++ {
        h.insertKey(i, 0, heightMap[i][0])
        h.insertKey(i, len(heightMap[0])-1, heightMap[i][len(heightMap[0]-1)])
        visited[i][0]
        visited[i][len(heightMap[0]-1)]
    }
    dir := [4]int{Dir{0,1}, Dir{1,0}, Dir{0,-1}, Dir{-1,0}}
    for len(h.arr) > 0 {
        currCell = h.extractMin()
        for i, _ := range dir {
            
        }
        
    }
    
    return water
}
type Dir struct {
    x int
    y int
}
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

func max(a int, b int) int {
	c := b
	if a > b {
		c = a
	}
	return c
}
