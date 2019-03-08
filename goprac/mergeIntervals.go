/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
/*
naive


*/
func merge(intervals []Interval) []Interval {
    mergeSort(&intervals, 0, len(intervals)-1)
    s := newStack()
    mergedIntervals := make([]Interval, 0)
    fmt.Println(intervals)
    for i := len(intervals)-1; i >= 0; i-- {
        s.push(&intervals[i])
    }
    
    for !s.isEmpty() {
        curr := s.pop()
        if doOverLap(curr,s.top()) {
            second := s.pop()
            s.push(mergeInterval(curr, second))
            
        } else {
             mergedIntervals = append(mergedIntervals, *curr)
        }
    }
    return mergedIntervals
}

type Stack struct {
    arr []*Interval
    size int
}

func newStack() Stack {
    var s Stack
    s.arr = make([]*Interval, 0)
    s.size = 0
    return s
}

func (s* Stack) push(k *Interval) {
    s.arr = append(s.arr, k)
    s.size++
}

func (s* Stack) pop() *Interval {
    if s.size > 0 {
        ret := s.arr[s.size-1]
        s.size--
        s.arr = s.arr[:s.size]
        return ret
    } 
    return nil
    
}

func (s *Stack) top() *Interval {
    if s.size > 0 {
        return s.arr[s.size-1]
    } 
    return nil
}

func (s* Stack) isEmpty() bool {
    return s.size == 0
}

func mergeSort(arr *[]Interval, start int, end int) {
    
    if start < end {
        mid := start + (end-start)/2
        
        mergeSort(arr, start, mid)
        mergeSort(arr, mid+1, end)
        
        mergeArr(arr, start, mid, end)
    }
    
}

func mergeArr(arr *[]Interval, start int, mid int, end int) {
    
    var indexL int
    var indexR int
    var indexA int
    n1 := mid - start + 1
    n2 :=  end - mid
    
    L := make([]Interval, n1)
    R := make([]Interval, n2)
    
    for i, _ := range L {
        L[i] = (*arr)[start+i]
    }
    
     for i, _ := range R {
        R[i] = (*arr)[mid+1+i]
    }
    
    indexL = 0
    indexR = 0
    indexA = start
    
    //merging into arr
    for indexL < n1 && indexR < n2 {
        if L[indexL].Start <= R[indexR].Start {
            (*arr)[indexA] = L[indexL]
            indexL++
        } else {
             (*arr)[indexA] = R[indexR]
            indexR++
        }
        indexA++
    }
    
    for indexL < n1 {
         (*arr)[indexA] = L[indexL]
        indexL++
        indexA++
    }
    for indexR < n2 {
         (*arr)[indexA] = R[indexR]
        indexR++
        indexA++
    }
    
}

func doOverLap(a *Interval, b *Interval) bool {
    if b == nil || a == nil {
        return false
    }
	return a.Start <= b.End && b.Start <= a.End
}

func mergeInterval(a *Interval, b *Interval) *Interval {
	return &Interval{min(a.Start, b.Start), max(a.End, b.End)}
}

func min(a int, b int) int {
	c := b
	if a < b {
		c = a
	}
	return c
}

func max(a int, b int) int {
	c := b
	if a > b {
		c = a
	}
	return c
}
