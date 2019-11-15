/*
prompt:

Partition to K Equal Sum Subsets

Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.

 

Example 1:

Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
Output: True
Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.
 
 
 
idea : use dfs to traverse the tree
however you must prune the search spaces
scenarios to consider: 
check if possible to divide total sum into partitions
check if biggest element can fit in one sack
if the element is equal to the size of the sack, then we can remove that sack and element from our search space

sorting the nums from greatest to smallest is faster because we fill up more buckets faster, therefore
when we get to the smaller numbers there are smaller amount of permutations the small numbers can be in
whereas if we start with the smaller numbers with more buckets, we have a bigger search space 

*/
func canPartitionKSubsets(nums []int, k int) bool {
    total := TotalSum(nums)
     //check if possible to divide total sum into partitions
    if total % k != 0 {
        return false
    }
    
    part := total / k
    sacks := make([]int,k)
     //check if biggest element can fit in one sack
    sort.Slice(nums, func(i, j int) bool {
        return nums[i]>nums[j]
    })
    if nums[0] > part {
        return false
    }
    sacks[0] += nums[0]
    //fmt.Println(part)
    //if the element is equal to the size of the sack, then we can remove that sack and element from our search space
    r := len(nums)-1
    for ; r >= 0 && nums[r] == part;  r-- {
        k--
    }
    if r < 0 {
        return true
    }
    return dfs(sacks[:k], nums[1:], k , part)
}

func dfs(sacks []int, nums[]int, k, part int) bool {
    allFull := true
    for i := range sacks {
        if sacks[i] != part {
            allFull = false
        }
    }
   // fmt.Println(sacks)
    if len(nums) == 0 || allFull {
        
        return allFull
    }
    
    for i,_ := range sacks {
        if sum := sacks[i] + nums[0]; sum <= part {
            sacks[i] = sum
            if dfs(sacks, nums[1:], k, part) {
                return true
            }
            sacks[i] -= nums[0]
            //if I take the element out and the sack is empty, then we know that we put this element in the other sacks
            //since all the sacks are empty after this one, since in dfs we fill the sacks from left to right
            if sacks[i] == 0 {
                break
            }
        }
    }
    return false
   
}


func TotalSum(nums []int) int {
    acc := 0 
    for i,_ := range nums {
        acc += nums[i]
    }
    return acc
}

