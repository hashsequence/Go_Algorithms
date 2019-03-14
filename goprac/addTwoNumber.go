/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var res *ListNode
    var resPtr *ListNode
    _ = resPtr
    ptr1 := l1
    ptr2 := l2
    carry := 0
    val := 0
    for ptr1 != nil || ptr2 != nil {
        val = ifNull(ptr1)+ifNull(ptr2) + carry
        carry = 0
        if val > 10 {
            val = val % 10 + carry
            carry = 1
            }
        if val == 10 {
            val = carry
            carry = 1
        }
        if res == nil {
            res = &ListNode{val, nil}
            resPtr = res
        } else {
            newNode := &ListNode{val,nil}
            resPtr.Next = newNode
            resPtr = resPtr.Next
        }
        incrementPtr(&ptr1,&ptr2)
    }
    if carry == 1{
        newNode := &ListNode{1,nil}
        resPtr.Next = newNode
        resPtr = resPtr.Next
    }
    return res
}

func ifNull(i *ListNode) int {
    if i == nil {
        return 0
    }
    return i.Val
}

func incrementPtr(ptr1, ptr2 **ListNode) {
    if *ptr1 != nil {
        *ptr1 = (*ptr1).Next
    }
    if *ptr2 != nil {
        *ptr2 = (*ptr2).Next
    }
}
