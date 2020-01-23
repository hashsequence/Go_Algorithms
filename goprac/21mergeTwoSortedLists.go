/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    //case where l1 is empty, l2 is empty, or both empty
    if l1 == nil && l2 == nil {
        return nil
    } else if l1 != nil && l2 == nil {
        return l1
    } else if l1 == nil && l2 != nil {
        return l2
    }
    //initialization 
    ptr1 := l1
    ptr2 := l2
    var head *ListNode
    var curr *ListNode
    if ptr1.Val < ptr2.Val {
        head = ptr1
        curr = ptr1
        ptr1 = ptr1.Next 
    } else if ptr1.Val > ptr2.Val {
        head = ptr2
        curr = ptr2
        ptr2 = ptr2.Next
    } else {
         head = ptr1
         curr = ptr1
         next1 := ptr1.Next
         next2 := ptr2.Next
         curr.Next = ptr2
         curr =  curr.Next
         ptr1 = next1
         ptr2 = next2
    }
    
    //iterate through list
    for ;ptr1 != nil && ptr2 != nil; {
        //fmt.Println(curr.Val)
        if ptr1.Val < ptr2.Val {
            next := ptr1.Next
            curr.Next = ptr1
            ptr1 = next
            curr = curr.Next
        } else if ptr1.Val > ptr2.Val {
            next := ptr2.Next
            curr.Next = ptr2
            ptr2 = next
            curr = curr.Next
        } else {
            next1 := ptr1.Next
            next2 := ptr2.Next
            curr.Next = ptr1
            curr = curr.Next
            curr.Next = ptr2
            curr = curr.Next
            ptr1 = next1
            ptr2 = next2
        }
   
        
    }
    //if lists are uneven append the rest of the longer list to the new list
    if curr != nil {
        if ptr1 == nil {
            curr.Next = ptr2
        } else {
            curr.Next = ptr1
        }
    }
    return head
}
