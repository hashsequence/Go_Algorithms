/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return head
    }
    if head.Next == nil {
        return head
    }
    var lessThanX *ListNode 
    var greaterThanX *ListNode
    var greaterThanXHead *ListNode
    var newHead *ListNode
    for it := head; it != nil; it = it.Next {
        if it.Val < x {
            if lessThanX == nil {
                lessThanX = &ListNode{it.Val,nil}
            } else {
                lessThanX.Next = &ListNode{it.Val,nil}
                lessThanX = lessThanX.Next
            }
            if newHead == nil {
                newHead = lessThanX
                continue
            }
            if newHead.Val >= x {
                 newHead = lessThanX
            }
        } else {
            if greaterThanX == nil {
                greaterThanX = &ListNode{it.Val,nil}
                greaterThanXHead =  greaterThanX
            } else {
                greaterThanX.Next = &ListNode{it.Val,nil}
                greaterThanX = greaterThanX.Next
            }
            if newHead == nil {
                newHead = greaterThanX
            }
        }
    }
    if lessThanX != nil {
        lessThanX.Next = greaterThanXHead
    }
    return newHead
}

