func fourSum(nums []int, target int) [][]int {
    res := [][]int{}
    
    //sort the nums 
    sort.Slice(nums, func(i,j int) bool {return nums[i] < nums[j]})
    fmt.Println(nums)
    //we iterate through the first pairs by using i and j
    //the 2nd pair we will iterate using two index from left and right
    //if the sum of 2 pairs is too great then we decrement the right index to decrease the value and 
    //if the sum of the 2 pair is too small then we increment the left index to increase value
    //to skip duplicates we dont need to check the range between l and r if l or r is the same as the prevous because we are checking the same
    //range with the same value
    //for i and j we dont need to check the same [l,r] since we considered all the values already
    for i := 0; i < len(nums)-3; i++ {
        for j := i+1; j < len(nums)-2; j++ {
           
            l := j+1
            r := len(nums)-1
            
            for l < r {
              //  fmt.Println(nums[i],nums[j],nums[l],nums[r])
                if nums[i] + nums[j] + nums[l] + nums[r] == target {
                    fourArr := make([]int,4)
                    fourArr[0], fourArr[1],fourArr[2],fourArr[3] = nums[i],nums[j],nums[l],nums[r]
                    res = append(res, fourArr)
                    for l < r && nums[l] == fourArr[2] {
                        l++
                    }
                    for l < r && nums[r] == fourArr[3]  {
                        r--
                    }
                    
                } else if nums[i] + nums[j] + nums[l] + nums[r] < target { 
                    l++
                } else {
                    r--
                }
            }
            for j < len(nums)-1 && nums[j] == nums[j+1] {
                j++
            }
        }
        for i < len(nums)-1 && nums[i] == nums[i+1] {
                i++
            }
    }
      return res
}
