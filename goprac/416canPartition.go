/*
problem:
Given a non-empty array containing only positive integers, find if the array
can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Note:

Each of the array element will not exceed 100.
The array size will not exceed 200.

ideas : two methods 
using dfs and recursion after sorting the slice from greatest to smallest 

second method:

use a set and iterate bfs 
get the halfsum
insert into the set
iterate through nums
for each num
we iterate through the current set and take the difference if the diff is positive
and insert into the set
for the next num in nums, if I see that num in the set again, then that means its partitionable since
that means the rest of the numbers are in the other set and the the numbers used to construct the diff is 
also in the set already, and the num we have now is in one of those sets
*/

func canPartition(nums []int) bool {
    if len(nums) <= 1 {
        return false
    }
    total := TotalSum(nums)
    if total % 2 != 0 {
        return false
    }
    half := total /2
    sort.Slice(nums,func(i,j int) bool {return nums[i] > nums[j]})
    return dfs(nums[1:], half - nums[0])
    
}

func dfs(nums []int, target int) bool {
    if target == 0 {
        return true
    } else if target < 0 {
        return false
    } 
    if len(nums) == 0 {
        return false
    }
    return dfs(nums[1:],target-nums[0]) || dfs(nums[1:],target)
}



func canPartitionv1(nums []int) bool {
    total := TotalSum(nums)
    if total % 2 != 0 {
        return false
    }
    half := total /2
    s := Set{}
    s.Insert(half)
    for i,_ := range nums {
        if !s.Exists(nums[i]) {
            copyS := Set{}
            for k, v := range s {
                copyS[k] = v
            }
            for keys,_ := range copyS {
                if keys > nums[i] {
                    s.Insert(keys-nums[i])
                }
            }
        } else {
            return true
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

type Set map[int]bool

func (s Set) Insert(k int) {
    s[k] = true
}

func (s Set) Remove(k int) {
    delete(s,k)
}

func (s Set) Exists(k int) bool {
    return s[k] 
}

