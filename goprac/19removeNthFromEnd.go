/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    i := 0
    if head == nil {
        return head
    }
    var nthNodeFromEnd *ListNode
    var prev *ListNode
    for curr := head; curr != nil; curr = curr.Next {
        if i == n-1 {
            nthNodeFromEnd = head
        } else if i >= n {
            prev = nthNodeFromEnd
            nthNodeFromEnd = nthNodeFromEnd.Next
        } 
        i++
    }
    if prev != nil {
        prev.Next = nthNodeFromEnd.Next
    }
    if nthNodeFromEnd == head {
        head = nthNodeFromEnd.Next
    }
    return head
}
