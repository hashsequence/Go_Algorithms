/*
Following is 2-bit sequence (n = 2)
  00 01 11 10
Following is 3-bit sequence (n = 3)
  000 001 011 010 110 111 101 100
And Following is 4-bit sequence (n = 4)
  0000 0001 0011 0010 0110 0111 0101 0100 1100 1101 1111 
  1110 1010 1011 1001 1000
  
basically the problem is to generate the gray sequence with n-bits

where the next element in the squence is a bit flipped

the algorithm is:

base case [0] if n == 0
base case [0,1] if n== 1
n >= 2
iterate i from 2 to n
    iterate the elements (e) you already have in your sequence backwards
        for element e you will make the ith msb a one and add this new element to your sequence
        eg. if e is 010 then add 1 to the left msb to become 110 


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

