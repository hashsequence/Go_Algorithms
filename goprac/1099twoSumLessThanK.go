/*
1099. Two Sum Less Than K
Easy


Given an array A of integers and integer K, return the maximum S such that there exists i < j with A[i] + A[j] = S and S < K. If no i, j exist satisfying this equation, return -1.

 

Example 1:

Input: A = [34,23,1,24,75,33,54,8], K = 60
Output: 58
Explanation: 
We can use 34 and 24 to sum 58 which is less than 60.
Example 2:

Input: A = [10,20,30], K = 15
Output: -1
Explanation: 
In this case it's not possible to get a pair sum less that 15.
 

Note:

1 <= A.length <= 100
1 <= A[i] <= 1000
1 <= K <= 2000

Solution:
sort the A
use window size r and l and shrink
according to sum
edge cases
if distanace of current sum is less than the stored closest but the current sum is > K then that cant be the closest
base case when length of array is 1 or 0

*/

func twoSumLessThanK(A []int, K int) int {
    if len(A) <= 1 {
        return -1
    }
    
    qSort(A)
    //fmt.Println(A)
    l := 0
    r := len(A)-1
    closest := A[l] + A[r]
    for l < r {
        if dist(A[l] + A[r], K) != 0 && dist(A[l] + A[r], K) <= dist(K, closest) {
            if  A[l] + A[r] < K {
                closest = A[l] + A[r]
            }  
            //fmt.Println(closest,A[l],A[r])
        }
        
        if K > A[l] + A[r] {
           l++
        } else {
           r--
        }
        
    }
    
    if closest > K {
        return -1
    }
    return closest
    
}

func dist(a int, b int) int {
    if a > b {
        return a-b
    } 
    return b-a
}
func qSort(a []int) []int {
  if len(a) < 2 { return a }

  left, right := 0, len(a) - 1

  // Pick a pivot
  pivotIndex := (right-left)/2

  // Move the pivot to the right
  a[pivotIndex], a[right] = a[right], a[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range a {
    if a[i] < a[right] {
      a[i], a[left] = a[left], a[i]
      left++
    }
  }

  // Place the pivot after the last smaller element
  a[left], a[right] = a[right], a[left]

  // Go down the rabbit hole
  qSort(a[:left])
  qSort(a[left + 1:])


  return a
}

