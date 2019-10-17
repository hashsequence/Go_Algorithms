func numTrees(n int) int {
  /*
  # of unique bst = catalan number
  C_n = Sigma i = 0 to i = n-1 C_i*C_(n-i-1)
  let F(n) := # of unique bst
  let G(i) := # of unique bst with ith node as the root out of n unique nodes
  G(i) = F_i*F_(n-i)
  base case: F_0 = 1
  */
    F := make([]int,n+1)
    F[0] = 1
    for i := 1; i < n+1; i++ {
        for j := 0; j < i; j++ {
            F[i] += F[j]*F[i-j-1] 
        }
    }
    return F[n]
}
