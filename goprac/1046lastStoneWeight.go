/*
1046. Last Stone Weight
Easy

We have a collection of rocks, each rock has a positive integer weight.

Each turn, we choose the two heaviest rocks and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:

If x == y, both stones are totally destroyed;
If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)

solution:
use priority queue
*/

func lastStoneWeight(stones []int) int {
    pq := make(PriorityQueue, len(stones))
    i := 0
    for _, val := range stones {
        pq[i] = &Item{
            value : val,
            priority: val,
            index : i,
        }
        i++
    }
    heap.Init(&pq)
  for pq.Len() >= 2 {
	  item1 := heap.Pop(&pq).(*Item)
      item2 := heap.Pop(&pq).(*Item)
      fmt.Println(item1.value, item2.value)
      if item1.value > item2.value {
          newItem := &Item{
               value : item1.value-item2.value,
               priority : item1.value-item2.value,
           }
          pq.Push(newItem)
      }
	}
    if pq.Len() == 0 {
        return 0
    }
    lastStone := pq.Pop().(*Item)
    return lastStone.value
}


type Item struct {
    value int
    priority int 
    index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {return len(pq)}

func (pq PriorityQueue) Less(i,j int) bool {
    return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq,item)
     pq.Update(item, item.value, item.priority)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    item.index = -1
    *pq = old[0 : n-1]
    return item
}

func (pq *PriorityQueue) Update(item *Item, value int, priority int) {
    item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
