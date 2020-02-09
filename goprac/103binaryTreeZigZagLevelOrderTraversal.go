/*
103. Binary Tree Zigzag Level Order Traversal
Medium

1516

84

Add to List

Share
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]

solution:
use two stacks so you can traverse level by level and alternate

to zigzag everytime you go another level switch order of inserting left and right nodes in current stack
start with inserting right then left

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    currStack := make(Stack, 0)
    otherStack := make(Stack, 0)
    leftOrRight := true //true is right, false is left 
    currStack.Push(root)
    res := [][]int{}
    
    for len(currStack) > 0 {
        levelArr := []int{}
        for len(currStack) > 0 {
            currNode := currStack.Pop()
            levelArr = append(levelArr, currNode.Val)
            if leftOrRight {
                if currNode != nil {
                    if currNode.Left != nil {
                        otherStack.Push(currNode.Left)
                    }
                    if currNode.Right != nil {
                        otherStack.Push(currNode.Right)
                    }
                }
            } else {
                 if currNode != nil {
                    if currNode.Right != nil {
                        otherStack.Push(currNode.Right)
                    }
                    if currNode.Left != nil {
                        otherStack.Push(currNode.Left)
                    }
                }
            }
        }
        res = append(res, levelArr)
        currStack, otherStack = otherStack, currStack
        leftOrRight = !leftOrRight
    }
    return res
}

type Stack []*TreeNode 

func (s *Stack) Top() *TreeNode {
    if len(*s) == 0 {
        return nil 
    }
    return (*s)[len(*s)-1]
}

func (s* Stack) Pop() *TreeNode {
    if len(*s) == 0 {
        return nil 
    }
    poppedNode := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return poppedNode
}

func (s* Stack) Push(node *TreeNode) {
    *s = append(*s, node)
}
