/*
572. Subtree of Another Tree
Easy


Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.

Example 1:
Given tree s:

     3
    / \
   4   5
  / \
 1   2
Given tree t:
   4 
  / \
 1   2
Return true, because t has the same structure and node values with a subtree of s.
Example 2:
Given tree s:

     3
    / \
   4   5
  / \
 1   2
    /
   0
Given tree t:
   4
  / \
 1   2
Return false.

solution:
must first look for cases
1st: s == null && t == null then true since nodes have reached the same spots
2nd: s is null and t is not null and vice versa, nodes have reached the same spot so false
3rd: check if t and s is the same tree if the value at current nodes are the same
4th check if t is a subtree of s.Left or s.right
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    } 
    if (s == nil && t != nil) || (s != nil && t == nil) {
        return false
    }
    if isBranchSame(s, t) && s.Val == t.Val {
        return true
    } 
    
    if isSubtree(s.Left, t) || isSubtree(s.Right,t) {
         return true
        
    }
    return false
}

func isBranchSame(s *TreeNode, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    } 
    if (s == nil && t != nil) || (s != nil && t == nil) {
        return false
    }
    if s.Val == t.Val {
        return isBranchSame(s.Left, t.Left) && isBranchSame(s.Right, t.Right)
    } else {
        return false
    }
}
