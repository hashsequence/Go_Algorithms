/*
114. Flatten Binary Tree to Linked List
Medium

Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:

    1
   / \
  2   5
 / \   \
3   4   6
The flattened tree should look like:

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 
 solution :
 use a stack to do inorder traversal
 remember to do base case if root is null
 */
func flatten(root *TreeNode)  {
    s := make(Stack,0)
    s.Push(root)
    for len(s) > 0 {
        curr := s.Front()
        s.Pop()
        if curr != nil && curr.Right != nil {
            s.Push(curr.Right)
        }
        if curr != nil && curr.Left != nil {
            s.Push(curr.Left)
        } 
        if len(s) > 0 {
            nextNode := s.Front()
            curr.Left = nil
            //fmt.Println("connecting ", curr.Val, " to ", nextNode.Val)
            curr.Right = nextNode
        }
    }
    
}

type Stack []*TreeNode

func (s *Stack) Front() *TreeNode {
    if len(*s) > 0 {
        return (*s)[len(*s)-1]
    }
    return nil
}

func (s *Stack) Push(val *TreeNode) {
    (*s) = append(*s,val)
} 

func (s *Stack) Pop() {
    if len(*s) == 1 {
        (*s)[0] = nil
        (*s) = make([]*TreeNode,0)
    }
    if len(*s) > 1 {
        (*s)[len(*s)-1] = nil
        (*s) = (*s)[:len(*s)-1]
    }
}

