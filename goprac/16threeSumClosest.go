/*
16. 3Sum Closest
Medium

Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

solution:
sort it first
base case if nums less or equal to three
n^2 algorithm for the loops
have l and r shrinking window type algorithm, decrement and increment r and l respectively, if curr sum is too small or too big
store curr sum if getting closer to target
imediately return if curr sum is equal to target

*/

func threeSumClosest(nums []int, target int) int {
    if len(nums) <= 3 {
        sum := 0
        for i,_ := range nums {
            sum += nums[i]
        }
        return sum
    }
    
    sort.Ints(nums)
    closest := nums[0] + nums[1] + nums[2]
    
    for i, _ := range nums {
        l := i+1
        r := len(nums)-1
        for ;l < r; {
            if distance(nums[i] + nums[l] + nums[r], target) == 0 {
                return nums[i] + nums[l] + nums[r]
            }
            if distance(nums[i] + nums[l] + nums[r], target) < distance(closest, target) {
                closest = nums[i] + nums[l] + nums[r]            }
            if nums[i] + nums[l] + nums[r] > target {
                r--
            } else {
                l++
            }
        }
    }
    return closest
}

func distance(a, b int) int {
    if a > b {
        return a - b
    } 
    return b - a
}
