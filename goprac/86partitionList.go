/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
the hard part was the many edge cases and knowing what to keep track of
we need a iterator pointer for the list that was less than x
we need a iterator pointer for the list that was greater or equal to x
we need a pointer to the head to the list that was greater or equal to the x so we can merge the two lists in the end
we also need to keep a pointer to the new head of the list

first thing is to decide what the new head was:
the new head will always be the head of the first node that is less than x,
but if we havent encountered a node less than x then we pick the first node that is greater than x

we then iterate over our original list with the above invariance


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

