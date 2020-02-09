/*
1. Two Sum
Easy

13414

490

Add to List

Share
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

solution:
remember to use a map
to map number to index
*/

func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, _ := range nums {
        //check if i've seen this number before, if not store target - current number 
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = i
        //check if target-nums[i] have been seen but is not the same as the current index
        } 
        if _, ok := m[target-nums[i]]; ok && m[target-nums[i]] != i {
            return []int{m[target-nums[i]], i}
        } 
    }
    return []int{}
}
