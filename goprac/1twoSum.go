func twoSum(nums []int, target int) []int {
    set := map[int]bool{}
    indices := map[int]int{}
    res := make([]int, 2)
    
    for i, _ := range nums {
            if len(set) == 0 {
                set[nums[i]] = true
                indices[nums[i]] = i
            } else {
                if set[target-nums[i]] == true {
                    res[0] = indices[target-nums[i]]
                    res[1] = i
                    return res
                } else {
                    set[nums[i]] = true
                    indices[nums[i]] = i
                }
            }
        }
   // fmt.Println(set)
   // fmt.Println(indices)
    return []int{}
}
