/*
Following is 2-bit sequence (n = 2)
  00 01 11 10
Following is 3-bit sequence (n = 3)
  000 001 011 010 110 111 101 100
And Following is 4-bit sequence (n = 4)
  0000 0001 0011 0010 0110 0111 0101 0100 1100 1101 1111 
  1110 1010 1011 1001 1000
  

the key idea is to find the gray sequence for n take the n-1 gray sequence add 0 the left msb and that gives you
the first half of the n gray sequence then add 1 to the left msb for the second half but add second half in reverse
ex graycode(2) =   00 01 11 10
graycode(3) = first half 000, 001, 011, 010   second half 110 111 101 100
*/

func grayCode(n int) []int {
    
    if n == 0 {
        return []int{0}

    }
    dp := []int{0,1}
    
    for i := 2; i < n+1; i++ {
        for j := len(dp)-1; j >= 0; j-- {
            dp = append(dp,(1 << uint(i-1)) | dp[j])
        }
    }
    return dp
}

