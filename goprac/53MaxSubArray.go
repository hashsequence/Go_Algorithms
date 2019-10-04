func maxSubArray(nums []int) int {
    return maxSubArrayReg(nums)
}

const MaxInt = int(^uint(0) >> 1) 
const MinInt = -MaxInt - 1

func maxSubArrayReg(nums []int) int {
  
    if len(nums) == 1 {
        return nums[0]
    } 
    if len(nums) == 0 {
        return 0
    }
    max := nums[0]
    sum := 0

    for i, _ := range nums {
        if sum < 0 {
            sum = nums[i]
        } else {
            sum += nums[i]
        }
        max = Max(max,sum)
    }
    return max
}

func Max( arr ...int ) int {
    m := arr[0]
    for _, val := range arr {
        if val > m {
            m = val
        }
    }
    return m
}
