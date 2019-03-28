/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    it1 := head
    it2 := head.Next
    newHead := it2
    var prev *ListNode
    
    for it2 != nil {
         if prev != nil {
            prev.Next = it2
        }
        it1.Next = it2.Next
        it2.Next = it1
        prev = it1
        it1 = it1.Next
        if it1 != nil {
            it2 = it1.Next
        } else {
            it2 = nil
        }
    }
    return newHead
}
