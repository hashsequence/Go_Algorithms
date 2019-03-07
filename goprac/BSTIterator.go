/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Stack struct {
    arr []*TreeNode
    size int
}

func newStack() Stack {
    var s Stack
    s.arr = make([]*TreeNode, 0)
    s.size = 0
    return s
}

func (s* Stack) push(k *TreeNode) {
    s.arr = append(s.arr, k)
    s.size++
}

func (s* Stack) pop() *TreeNode {
    if s.size > 0 {
        ret := s.arr[s.size-1]
        s.size--
        s.arr = s.arr[:s.size]
        return ret
    } 
    return nil
    
}

func (s *Stack) top() *TreeNode {
    if s.size > 0 {
        return s.arr[s.size-1]
    } 
    return nil
}

func (s* Stack) isEmpty() bool {
    return s.size == 0
}

type BSTIterator struct {
    s Stack
    it *TreeNode
    hasNext bool
}

func getLeftMostTreeNode(root *TreeNode, s *Stack) *TreeNode{
    for i := root; i != nil; i = i.Left {
             s.push(i)
    } 
    return s.top()
}

func Constructor(root *TreeNode) BSTIterator {
    var t BSTIterator
    t.s = newStack()
    t.it = getLeftMostTreeNode(root, &t.s)
    if t.s.isEmpty() {
        t.hasNext = false
    } else {
        t.hasNext = true
    }
    return t
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    fmt.Println(this.s)
    next := this.s.pop()
    if next == nil {
        return this.it.Val
    } 
    this.it = getLeftMostTreeNode(next.Right, &this.s)
    if this.s.isEmpty() {
        this.hasNext = false
    }
    return next.Val
    
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return this.hasNext 
}



/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
