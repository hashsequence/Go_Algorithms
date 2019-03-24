/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 || head == nil || head.Next == nil {
        return head
    }
    
    it := head
    var itStartPrev *ListNode
    itStart := head
    i := 0
    first := true
    for  {
        if i == k {
           // fmt.Println(itStart.Val, " ", it.Val)
            temp := reverseLinkedList(itStart, it)
            if itStartPrev != nil {
               itStartPrev.Next = temp 
            }
            itStartPrev = itStart
            itStart = it 
             if first {
                head = temp
                first = false
            }
            i = 0
        }
        if it == nil {
            break
        }
        it = it.Next
        i++
    }
    return head
}

func reverseLinkedList(head *ListNode, tail *ListNode) (*ListNode) {
    
    if head == nil || head.Next == nil || head == tail {
        return head
    }
    
    it := head
    itNext := it.Next
    var itPrev *ListNode
    
    for it != tail {
        itNext = it.Next
        it.Next = itPrev
        itPrev = it
        it = itNext
    }
    head.Next = tail
    head = itPrev
    return head
}
