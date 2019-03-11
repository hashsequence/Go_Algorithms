/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 
 type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}
 */

func insert(intervals []Interval, newInterval Interval) []Interval {
    lenIntervals := len(intervals)
    if lenIntervals == 0 {
        return append(intervals, newInterval)
    }
    s := newStack()
    newIndex := 0
    insertedFlag := 0
    for  i := 0; i < lenIntervals; i++ {
        if less(newInterval, intervals[i]) {
            newIndex = i
            insertedFlag = 1
            break
        }
       
    }
    //this means the newInterval is greater than all the other interval in intervals
    if insertedFlag == 0 {
        newIndex = lenIntervals
        lenIntervals++
    }
    
     for j := lenIntervals-1; j >= 0; j-- {
         if j < len(intervals) {
              s.push(&intervals[j])
         }
         if j == newIndex {
             s.push(&newInterval)
         }
    }
    
    mergedIntervals := make([]Interval, 0)

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

/*
type IntervalList []Interval

func (this *IntervalList) Len() int {
    return len(this)
}

func (this *IntervalList) Less(i, j int) bool {
    return this[i].Start < this[j].Start
}

func (this *IntervalList) Swap(i, j int) {
    t := this[i]
    this[i] = this[j]
    this[j] = t
} 

func prepend(dest []Interval, value Interval) []Interval {
    if cap(dest) > len(dest) {
        dest = dest[:len(dest)+1]
        copy(dest[1:], dest)
        dest[0] = value
        return dest
    }

    // No room, new slice need to be allocated:
    // Use some extra space for future:
    res := make([]Interval, len(dest)+1, len(dest)+5)
    res[0] = value
    //func copy(dst, src []Type) int
    copy(res[1:], dest)
    return res
}
*/

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

func less(i, j Interval) bool {
    return i.Start < j.Start
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


