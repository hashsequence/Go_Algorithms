/*
41. First Missing Positive
Hard


Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.

solution:
the first thing you must notice that if we remove all elements less than or equal to zero
we will be left with n positive numbers,
and so the smallest positive number is in [1,n]

base case is if we dont see a one then the answer is 1
if we do see a 1 but the nums size is 1 then the answer is 2

we need a boolean arr[n] and as we look at each positive number less than or equal to n we 
mark arr[n-1] = true
and so the biggest index where arr[n-1] is false is the answer (n-1 + 1 = n) or if its all true
then n+1 is the answer
we can use the nums array as a boolean array by flipping the sign for a[i], -1 means i is seen and > 0 means 
i is not in nums
*/

func firstMissingPositive(nums []int) int {
    n := 0
    seenOne := false
    //eliminate all numbers less than or equal to 0 by changing it to 1, but before we do we want to see if it was
    //a one before
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            seenOne = true
        }
        if nums[i] <= 0 {
            nums[i] = 1
        } else {
            n++
        }
    }
    //now we have n which is the number of positive numbers
    //here are base cases
    if seenOne == false {
        return 1
    }
    if seenOne == true && n == 1 {
        return 2
    }
    //we flip the sign (but only once) at position i of the array if i is in nums
    for i := 0; i < len(nums); i++ {
        if Abs(nums[i]) <= n && nums[Abs(nums[i])-1] > 0 {
            nums[Abs(nums[i])-1] *= -1
        }
    }
    //fmt.Println(n, nums)
    //find the first nonnegative and the ith positive is the smallest positive
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            return i+1
        }
    }
    //all numbers where seen in the range [1,n] so the answer is n+1
    return n+1
    
}

func Abs (i int) int {
    if i < 0 {
        return -i
    }
    return i
}
