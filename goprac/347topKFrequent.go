func topKFrequent(nums []int, k int) []int {
    countArr := map[int]int{}
    
    for i, _ := range nums {
        countArr[nums[i]] += 1
    }
    
    bucket := make([]*[]int, len(nums))
    for i, _ := range countArr {
        if bucket[len(nums)-countArr[i]] == nil {
            bucket[len(nums)-countArr[i]] = new([]int)
        }
        *bucket[len(nums)-countArr[i]] = append(*bucket[len(nums)-countArr[i]], i)
    }
    
    res := make([]int,k)
    counter := 0
    for i, _ := range bucket {
        if bucket[i] != nil {
            for j,_ := range *bucket[i] {
                if counter == k {
                    break
                }
                res[counter] = (*bucket[i])[j]
                counter++
            }
        }
    }
    return res
}
