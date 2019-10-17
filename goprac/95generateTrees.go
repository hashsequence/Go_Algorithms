/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*

the idea is that you can divide into further subproblems
let F(n) := be the set of all unique bst
let G(i,n) := be the set of all unique bst with root node i out on n nodes

so for root i , all the lefttree nodes is with bsts with roots in [0,...,i-1] and righttree nodes with roots in [i+1,...,n]

so let basically G(i,n) := [unique bst with root i with with roots in [0,...,i-1] and righttree nodes with roots in [i+1,...,n]]

lets make a function called generateTreesFromSet that takes in a set of unqiue values and return a set of unique bst with those values
so visually

              i
              *
             * *  
            *   *
   gt[0:i-1]    gt[i+1:n]

*/
func generateTrees(n int) []*TreeNode {
    if n <= 0 {
        return []*TreeNode{}
    }
    set := make([]int,n)
    for i := 1 ; i < n+1; i++ {
        set[i-1] = i
    }
    return generateTreesFromSet(set)
}

func generateTreesFromSet(set []int) []*TreeNode {
    if len(set) == 0 {
        return []*TreeNode{nil}
    }
    if len(set) == 1 {
        return []*TreeNode{&TreeNode{set[0],nil,nil}}
    }
    arrOfTrees := []*TreeNode{}
    for i := 0; i < len(set); i++ {
        leftTrees := generateTreesFromSet(set[0:i])
        rightTrees := generateTreesFromSet(set[i+1:])
        root := set[i]
        for k,_ := range leftTrees {
            for l,_ := range rightTrees{
                arrOfTrees = append(arrOfTrees, &TreeNode{root,leftTrees[k],rightTrees[l]})
            }
        }
    }
    return arrOfTrees
}
